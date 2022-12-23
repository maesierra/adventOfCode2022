package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

var logLevel int = 0

var rotations map[string]map[string]string = map[string]map[string]string{
	"L": {
		"^": "<",
		">": "^",
		"v": ">",
		"<": "v",

	},
	"R": {
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",

	},
	"O": {
		"^": "v",
		">": "<",
		"v": "^",
		"<": ">",

	},
}

type Day22 struct {
}

type Tile struct {
	row int
	column int
	tileType string
	up *Tile
	upRotation string
	right *Tile
	rightRotation string
	down *Tile
	downRotation string
	left *Tile
	leftRotation string
}

func (t Tile) String() string {
	return t.Key()
}
func (t Tile) Key() string {
	return fmt.Sprintf("%v-%v", t.row, t.column)
}

func (t Tile) Next(direction string) (*Tile,string) {
	switch direction {
	case ">":
		return t.right, t.rightRotation
	case "<":
		return t.left, t.leftRotation
	case "^":
		return t.up, t.upRotation
	case "v":
		return t.down, t.downRotation
	}
	panic("unknown direction")
}

type Trail struct {
	position common.Point
	direction string
}
type Board struct {
	tiles map[string]*Tile
	rows int
	columns int
	Position common.Point
	Direction string
	trail []Trail
}


func (b Board) Tile(row, column int) *Tile {
	t, present := b.tiles[fmt.Sprintf("%v-%v", row, column)]
	if !present {
		return nil
	} else {
		return t
	}
}

func (b Board) TileForPoint(p common.Point) *Tile {
	return b.Tile(p.Y, p.X)
}


func (b Board) String() string {
	return b.Print(false)
}

func (b Board) Print(trail bool) string {
	trailPoints := map[string]string{}
	if trail {
		for _, t := range b.trail {
			key := fmt.Sprintf("%v-%v", t.position.Y, t.position.X)
			trailPoints[key] = t.direction
		}	
	}
	str := ""
	for row := 0; row < b.rows; row++ {
		str += fmt.Sprintf("% 4d ", row)
		for column := 0; column < b.columns; column++ {
			if b.Position.X == column && b.Position.Y == row {
				str += b.Direction
				continue
			} 
			key := fmt.Sprintf("%v-%v", row, column)
			if trail, present := trailPoints[key]; present {
				str += trail
				continue
			}
			t := b.Tile(row, column)
			if t != nil {
				str += t.tileType
			} else {
				str += " "
			}
		}
		str += "\n"
	}
	return str
}
func (b *Board) Move(m Movement) {
	//Move in the curring direction until wall
	tile := b.Tile(b.Position.Y, b.Position.X)	
	prev := tile
	rotation := ""
	for n := 0; n < m.N; n++ {
		tile, rotation = tile.Next(b.Direction)		
		if tile.tileType == "#" {
			break
		}
		prev = tile
		if rotation != "" {
			b.Direction = rotations[rotation][b.Direction]
		}
		b.trail =append(b.trail, Trail{
			position: common.Point{X: tile.column, Y: tile.row},
			direction: b.Direction,
		})
	}
	b.Position.X = prev.column
	b.Position.Y = prev.row
	if m.Rotation != ""{
		//Rotate
		b.Direction = rotations[m.Rotation][b.Direction]
	}

}

type Movement struct {
	N int
	Rotation string
}


type Connection struct {
	face  int
	vertex string
	rotation string
}

type Face struct {
	id int
	row int
	column int 
	top Connection
	left Connection
	right Connection
	bottom Connection
}

var Maps2d = map[int]Map {
	1: {
		id: 1,
		faces: [6]Face{
			{id: 1, row: 0, column: 2, top: Connection{5, "B", ""}, left: Connection{1, "R", ""}, right: Connection{1, "L", ""}, bottom: Connection{4, "T", ""}},
			{id: 2, row: 1, column: 0, top: Connection{2, "B", ""}, left: Connection{4, "R", ""}, right: Connection{3, "L", ""}, bottom: Connection{2, "T", ""}},
			{id: 3, row: 1, column: 1, top: Connection{3, "B", ""}, left: Connection{2, "R", ""}, right: Connection{4, "L", ""}, bottom: Connection{2, "T", ""}},
			{id: 4, row: 1, column: 2, top: Connection{1, "B", ""}, left: Connection{3, "R", ""}, right: Connection{2, "L", ""}, bottom: Connection{5, "T", ""}},
			{id: 5, row: 2, column: 2, top: Connection{4, "B", ""}, left: Connection{6, "R", ""}, right: Connection{6, "L", ""}, bottom: Connection{1, "T", ""}},
			{id: 6, row: 2, column: 3, top: Connection{6, "B", ""}, left: Connection{5, "R", ""}, right: Connection{5, "L", ""}, bottom: Connection{6, "T", ""}},
		},
		dimensions: [2]int{3, 4},
	}, 
	2: {
		id: 2,
		faces: [6]Face{
			{id: 1, row: 0, column: 1, top: Connection{5, "B", ""}, left: Connection{6, "R", ""}, right: Connection{6, "L", ""}, bottom: Connection{4, "T", ""}},
			{id: 6, row: 0, column: 2, top: Connection{6, "B", ""}, left: Connection{1, "R", ""}, right: Connection{1, "L", ""}, bottom: Connection{6, "T", ""}},
			{id: 4, row: 1, column: 1, top: Connection{1, "B", ""}, left: Connection{4, "R", ""}, right: Connection{4, "L", ""}, bottom: Connection{5, "T", ""}},
			{id: 3, row: 2, column: 0, top: Connection{2, "B", ""}, left: Connection{5, "R", ""}, right: Connection{5, "L", ""}, bottom: Connection{2, "T", ""}},
			{id: 5, row: 2, column: 1, top: Connection{4, "B", ""}, left: Connection{3, "R", ""}, right: Connection{3, "L", ""}, bottom: Connection{1, "T", ""}},
			{id: 2, row: 3, column: 0, top: Connection{3, "B", ""}, left: Connection{2, "R", ""}, right: Connection{2, "L", ""}, bottom: Connection{3, "T", ""}},
		},
		dimensions: [2]int{4, 3},
	},

}

var Maps3d = map[int]Map {
	1: {
		id: 1,
		faces: [6]Face{
			{id: 1, row: 0, column: 2, top: Connection{2, "T", "O"}, left: Connection{3, "T", "L"}, right: Connection{6, "R", "O"}, bottom: Connection{4, "T", "" }},
			{id: 2, row: 1, column: 0, top: Connection{1, "T", "O"}, left: Connection{6, "B", "R"}, right: Connection{3, "L", "" }, bottom: Connection{5, "B", "O"}},
			{id: 3, row: 1, column: 1, top: Connection{1, "L", "R"}, left: Connection{2, "R", "" }, right: Connection{4, "L", "" }, bottom: Connection{5, "L", "L"}},
			{id: 4, row: 1, column: 2, top: Connection{1, "B", "" }, left: Connection{3, "R", "" }, right: Connection{6, "T", "R"}, bottom: Connection{5, "T", "" }},
			{id: 5, row: 2, column: 2, top: Connection{4, "B", "" }, left: Connection{3, "B", "R"}, right: Connection{6, "L", "" }, bottom: Connection{2, "B", "O"}},
			{id: 6, row: 2, column: 3, top: Connection{6, "R", "L"}, left: Connection{5, "R", "" }, right: Connection{1, "R", "O"}, bottom: Connection{2, "L", "L"}},
		},
		dimensions: [2]int{3, 4},
	}, 
	2: {
		id: 2,
		faces: [6]Face{
			{id: 1, row: 0, column: 1, top: Connection{2, "L", "R"}, left: Connection{3, "L", "O"}, right: Connection{6, "L", "" }, bottom: Connection{4, "T", "" }},
			{id: 2, row: 3, column: 0, top: Connection{3, "B", "" }, left: Connection{1, "T", "L"}, right: Connection{5, "B", "L"}, bottom: Connection{6, "T", ""}},
			{id: 3, row: 2, column: 0, top: Connection{4, "L", "R"}, left: Connection{1, "L", "O"}, right: Connection{5, "L", "" }, bottom: Connection{2, "T", "" }},
			{id: 4, row: 1, column: 1, top: Connection{1, "B", "" }, left: Connection{3, "T", "L"}, right: Connection{6, "B", "L"}, bottom: Connection{5, "T", "" }},
			{id: 5, row: 2, column: 1, top: Connection{4, "B", "" }, left: Connection{3, "R", "" }, right: Connection{6, "R", "O"}, bottom: Connection{2, "R", "R"}},
			{id: 6, row: 0, column: 2, top: Connection{2, "B", "" }, left: Connection{1, "R", "" }, right: Connection{5, "R", "O"}, bottom: Connection{4, "R", "R"}},

		},
		dimensions: [2]int{4, 3},
	},

}

type Map struct {
	id int
	faces [6]Face
	dimensions [2]int
	rows int
	columns int
}

func (m Map) FaceHeight() int {
	return m.rows / m.dimensions[0]
}

func (m Map) FaceWidth() int {
	return m.columns / m.dimensions[1]
}


func (m Map) GetFace(id int) Face {
	for _, f := range m.faces {
		if f.id == id {
			return f
		}
	}
	panic("no face")
}

func (m Map) ConnectTop(face Face, column int) common.Point {
	delta := column - face.column * m.FaceWidth() 
	transpose := false
	otherFace := m.GetFace(face.top.face)
	switch face.top.vertex {
	case "T": 
		column = (otherFace.column + 1) * m.FaceWidth() - 1 - delta
	case "L":
		transpose = true
		column = otherFace.row  * m.FaceHeight() +  delta		
	case "R": 
		transpose = true
		column = (otherFace.row + 1) * m.FaceHeight() - 1 - delta		
	case "B": 
		column = otherFace.column  * m.FaceWidth() +  delta		
	}
	for _, p := range m.GetVertex(otherFace, face.top.vertex) {
		if !transpose && p.X == column {
			return p
		} else if transpose && p.Y == column {
			return p
		}
	}
	panic("no connect top")
}

func (m Map) ConnectBottom(face Face, column int) common.Point {
	delta := column - face.column * m.FaceWidth() 
	transpose := false
	otherFace := m.GetFace(face.bottom.face)
	switch face.bottom.vertex {
	case "B": 
		column = (otherFace.column + 1) * m.FaceWidth() - 1 - delta
	case "L":
		transpose = true
		column = (otherFace.row + 1) * m.FaceHeight() - 1 - delta		
	case "R": 
		transpose = true
		column = otherFace.row  * m.FaceHeight() +  delta		
	case "T": 
		column = otherFace.column  * m.FaceWidth() +  delta		
	}
	for _, p := range m.GetVertex(otherFace, face.bottom.vertex) {
		if !transpose && p.X == column {
			return p
		} else if transpose && p.Y == column {
			return p
		}
	}
	panic("no connect bottom")
}

func (m Map) ConnectLeft(face Face, row int) common.Point {
	delta := row - face.row * m.FaceHeight() 
	transpose := false
	otherFace := m.GetFace(face.left.face)
	switch face.left.vertex {
	case "B": 
		transpose = true
		row = (otherFace.column + 1) * m.FaceWidth() - 1 - delta
	case "L":
		row = (otherFace.row + 1) * m.FaceHeight() - 1 - delta		
	case "R": 
		row = otherFace.row * m.FaceHeight() + delta
	case "T": 
		transpose = true
		row = otherFace.column * m.FaceHeight() + delta
	}
	for _, p := range m.GetVertex(otherFace, face.left.vertex) {
		if !transpose && p.Y == row {
			return p
		} else if transpose && p.X == row {
			return p
		}
	}
	panic("no connect left")
}

func (m Map) ConnectRight(face Face, row int) common.Point {
	delta := row - face.row * m.FaceHeight() 
	transpose := false
	otherFace := m.GetFace(face.right.face)
	switch face.right.vertex {
	case "B": 
		transpose = true
		row = otherFace.column * m.FaceWidth() + delta
	case "L":
		row = otherFace.row * m.FaceHeight() + delta
	case "R": 
		row = (otherFace.row + 1) * m.FaceHeight() - 1 - delta		
	case "T": 
		transpose = true
		row = (otherFace.column + 1) * m.FaceWidth() - 1 - delta
	}
	for _, p := range m.GetVertex(otherFace, face.right.vertex) {
		if !transpose && p.Y == row {
			return p
		} else if transpose && p.X == row {
			return p
		}
	}
	panic("no connect right")
}

func (m Map) GetVertex(face Face, vertex string) []common.Point {
	switch vertex {
	case "B":
		return m.GetBottomBorder(face)
	case "T":
		return m.GetTopBorder(face)
	case "R":
		return m.GetRightBorder(face)
	case "L":
		return m.GetLeftBorder(face)
	}
	panic("unknown vertex")
}

func (m Map) GetTopBorder(face Face) []common.Point {
	res := []common.Point{}
	from := face.column * m.FaceWidth()
	to := (face.column + 1) * m.FaceWidth()
	y := face.row * m.FaceHeight()
	for column := from; column < to; column++ {
		res = append(res, common.Point{X: column, Y: y})
	}
	return res
}

func (m Map) GetBottomBorder(face Face) []common.Point {
	res := []common.Point{}
	from := face.column * m.FaceWidth()
	to := (face.column + 1) * m.FaceWidth()
	y := ((face.row + 1) * m.FaceHeight()) - 1
	for column := from; column < to; column++ {
		res = append(res, common.Point{X: column, Y: y})
	}
	return res
}

func (m Map) GetLeftBorder(face Face) []common.Point {
	res := []common.Point{}
	from := face.row * m.FaceHeight()
	to := (face.row + 1) * m.FaceHeight()
	x := face.column * m.FaceWidth()
	for row := from; row < to; row++ {
		res = append(res, common.Point{X: x, Y: row})
	}
	return res
}

func (m Map) GetRightBorder(face Face) []common.Point {
	res := []common.Point{}
	from := face.row * m.FaceHeight()
	to := (face.row + 1) * m.FaceHeight()
	x := ((face.column + 1) * m.FaceWidth()) - 1
	for row := from; row < to; row++ {
		res = append(res, common.Point{X: x, Y: row})
	}
	return res
}



func (d Day22) ParseInput(inputFile string, m *Map) (Board, []Movement) {

	input := common.ReadFileIntoLBlocks(inputFile, "\n\n")
	blockLines := strings.Split(input[0], "\n")
	movementsLine := input[1]


	columns := 0
	//Complete all the lines to the same size avoid problems parsing later
	rows := len(blockLines)
	for _, line := range blockLines {
		columns = common.IntMax(columns, len(line))
	}
	for idx, line := range blockLines {
		if len(line) < columns {
			blockLines[idx] = line + strings.Repeat(" ", columns - len(line))
		}
	}
	m.rows = rows
	m.columns = columns
	board := Board{
		tiles: map[string]*Tile{},
		rows: rows,
		columns: columns,
		trail: []Trail{},
	}


	for row, line := range blockLines {
		for column, ch := range line {
			if ch == ' ' {
				continue
			}
			tile := Tile{
				row: row,
				column: column,
				tileType: string(ch),
			}
			board.tiles[tile.Key()] = &tile
			//Connect with the neighbours
			if column != 0 {
				t := board.Tile(row, column - 1)
				if t != nil {
					tile.left = t
					t.right = &tile
				} 
			} else if column != len(line) - 1  {
				t := board.Tile(row, column + 1)
				if t != nil {
					tile.right = t
					t.left = &tile
				}
			}
			if row != 0 {
				t := board.Tile(row - 1, column)
				if t != nil {
					tile.up = t
					t.down = &tile
				}
			} else if row != len(blockLines) - 1 {
				t := board.Tile(row + 1, column)
				if t != nil {
					tile.down = t
					t.up = &tile
				}
			}
		}
	}

	//Close the borders according to the map
	for _, face := range m.faces {
		for _, p := range m.GetTopBorder(face) {
			t := board.TileForPoint(p)
			if t.up == nil {
				otherPoint := m.ConnectTop(face, p.X)
				other := board.TileForPoint(otherPoint)				
				t.up = other
				t.upRotation = face.top.rotation
			}
		}
		for _, p := range m.GetRightBorder(face) {
			t := board.TileForPoint(p)
			if t.right == nil {
				otherPoint := m.ConnectRight(face, p.Y)
				other := board.TileForPoint(otherPoint)				
				t.right = other
				t.rightRotation = face.right.rotation
			}
		} 
		for _, p := range m.GetBottomBorder(face) {
			t := board.TileForPoint(p)
			if t.down == nil {
				otherPoint := m.ConnectBottom(face, p.X)
				other := board.TileForPoint(otherPoint)				
				t.down = other
				t.downRotation = face.bottom.rotation
			}
		}
		for _, p := range m.GetLeftBorder(face) {
			t := board.TileForPoint(p)
			if t.left == nil {
				otherPoint := m.ConnectLeft(face, p.Y)
				other := board.TileForPoint(otherPoint)				
				t.left = other
				t.leftRotation = face.left.rotation
			}
		}
	}
	
	//parse the movements
	r, _ := regexp.Compile(`(\d+)(R|L)?`)
	movements := []Movement{}
	for _, m := range r.FindAllStringSubmatch(movementsLine, -1) {
		n, _ := strconv.Atoi(m[1])
		movements = append(movements, Movement{N: n, Rotation: m[2]})
	}
	//set up the starting position
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			t := board.Tile(row, column)
			if t != nil && t.tileType != "#" {
				board.Position = common.Point{
					X: column,
					Y: row,
				}
				board.Direction = ">"
				board.trail = append(board.trail, Trail{position: board.Position, direction: board.Direction})
				return board, movements			
			}
		}
	}
	panic("no starting position")
}


func (d Day22) Run(inputFile string, m Map) int {
	board, movements := d.ParseInput(inputFile, &m)
	if logLevel > 0 {
		fmt.Printf("%v\n\n", board)
	}
	for idx, move := range movements {
		current := board.Direction
		board.Move(move)
		if logLevel > 2 {
			fmt.Printf("%v after move %v%v to %v\n%v\n\n", idx, move.N, move.Rotation, current, board)
		}
		if logLevel > 1 && idx%20 == 19 {
			fmt.Println(board.Print(true))
		}
		fmt.Printf("%v%v: %d,%d %v\n", move.N, move.Rotation, board.Position.Y+1, board.Position.X+1, board.Direction)
	}
	if logLevel > 1 {
		fmt.Println(board.Print(true))
	}
	fmt.Printf("final position %v %v\n", board.Position.Y+1, board.Position.X+1)
	var password int = ((board.Position.Y + 1) * 1000) + ((board.Position.X + 1) * 4)
	switch board.Direction {
	case ">":
		password += 0
	case "<":
		password += 2
	case "^":
		password += 3
	case "v":
		password += 1

	}
	return password
}

func (d Day22) SolvePart1(inputFile string, data []string) string {

	mapId := 2
	if len(data) > 0 {
		v, _ := strconv.Atoi(data[0])
		mapId = v
	}	

	m := Maps2d[mapId]
	password := d.Run(inputFile, m)
	return strconv.Itoa(int(password))

}

func (d Day22) SolvePart2(inputFile string, data []string) string {
	mapId := 2
	if len(data) > 0 {
		v, _ := strconv.Atoi(data[0])
		mapId = v
	}	

	m := Maps3d[mapId]
	password := d.Run(inputFile, m)
	return strconv.Itoa(int(password))
}
