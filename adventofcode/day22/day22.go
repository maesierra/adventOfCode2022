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
			b.Direction = rotations[m.Rotation][b.Direction]
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



type Face struct {
	id int
	row int
	column int 
	top int
	topRotation string
	left int
	leftRotation string
	right int
	rightRotation string
	bottom int
	bottomRotation string
}
var Maps2d = map[int]Map {
	1: {
		id: 1,
		faces: [6]Face{
			{id: 1, row: 0, column: 2, top: 5, left: 1, right: 1, bottom: 4},
			{id: 2, row: 1, column: 0, top: 2, left: 4, right: 3, bottom: 2},
			{id: 3, row: 1, column: 1, top: 3, left: 2, right: 4, bottom: 2},
			{id: 4, row: 1, column: 2, top: 1, left: 3, right: 2, bottom: 5},
			{id: 5, row: 2, column: 2, top: 4, left: 6, right: 6, bottom: 1},
			{id: 6, row: 2, column: 3, top: 6, left: 5, right: 5, bottom: 6},
		},
		dimensions: [2]int{3, 4},
	}, 
	2: {
		id: 2,
		faces: [6]Face{
			{id: 1, row: 0, column: 1, top: 5, left: 6, right: 6, bottom: 4},
			{id: 6, row: 0, column: 2, top: 6, left: 1, right: 1, bottom: 6},
			{id: 4, row: 1, column: 1, top: 1, left: 4, right: 4, bottom: 5},
			{id: 3, row: 2, column: 0, top: 2, left: 5, right: 5, bottom: 2},
			{id: 5, row: 2, column: 1, top: 4, left: 3, right: 3, bottom: 1},
			{id: 2, row: 3, column: 0, top: 3, left: 2, right: 2, bottom: 3},
		},
		dimensions: [2]int{4, 3},
	},

}

var Maps3d = map[int]Map {
	1: {
		id: 1,
		faces: [6]Face{
			{id: 1, row: 0, column: 2, top: 2, topRotation: "O", left: 3, leftRotation: "L", right: 6, rightRotation: "O", bottom: 4, bottomRotation: "" },
			{id: 2, row: 1, column: 0, top: 1, topRotation: "O", left: 6, leftRotation: "L", right: 3, rightRotation: "",  bottom: 5, bottomRotation: "O"},
			{id: 3, row: 1, column: 1, top: 1, topRotation: "R", left: 2, leftRotation: "",  right: 4, rightRotation: "",  bottom: 5, bottomRotation: "L"},
			{id: 4, row: 1, column: 2, top: 1, topRotation: "",  left: 3, leftRotation: "",  right: 6, rightRotation: "R", bottom: 5, bottomRotation: "" },
			{id: 5, row: 2, column: 2, top: 4, topRotation: "",  left: 3, leftRotation: "R", right: 6, rightRotation: "",  bottom: 2, bottomRotation: "O"},
			{id: 6, row: 2, column: 3, top: 6, topRotation: "L", left: 5, leftRotation: "",  right: 1, rightRotation: "O", bottom: 2, bottomRotation: "R"},
		},
		dimensions: [2]int{3, 4},
	}, 
	2: {
		id: 2,
		faces: [6]Face{
			{id: 1, row: 0, column: 1, top: 2, left: 3, right: 6, bottom: 4},
			{id: 6, row: 0, column: 2, top: 6, left: 5, right: 1, bottom: 2},
			{id: 4, row: 1, column: 1, top: 4, left: 3, right: 6, bottom: 2},
			{id: 5, row: 2, column: 1, top: 1, left: 3, right: 6, bottom: 5},
			{id: 3, row: 2, column: 0, top: 1, left: 2, right: 4, bottom: 5},
			{id: 2, row: 3, column: 0, top: 1, left: 6, right: 3, bottom: 5},
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

func (m Map) GetFace(row, column int) *Face  {
	faceRow := row / m.FaceHeight()
	faceColumn := column / m.FaceWidth()
	for _, f:= range m.faces {
		if f.row == faceRow && f.column == faceColumn {
			return &f
		} 
	}
	return nil
}

func (m Map) ConnectTop(face Face, column int) (*Face, common.Point) {
	for _, otherFace := range m.faces {
		if otherFace.id == face.top {
			for _, p := range m.GetBottomBorder(otherFace) {
				if p.X == column {
					return &otherFace, p
				}
			}
		}
	}
	panic("no connect top")
}

func (m Map) ConnectBottom(face Face, column int) (*Face, common.Point) {
	for _, otherFace := range m.faces {
		if otherFace.id == face.bottom {
			for _, p := range m.GetTopBorder(otherFace) {
				if p.X == column {
					return &otherFace, p
				}
			}
		}
	}
	panic("no connect bottom")
}

func (m Map) ConnectLeft(face Face, row int) (*Face, common.Point) {
	for _, otherFace := range m.faces {
		if otherFace.id == face.left {
			for _, p := range m.GetRightBorder(otherFace) {
				if p.Y == row {
					return &otherFace, p
				}
			}
		}
	}
	panic("no connect left")
}

func (m Map) ConnectRight(face Face, row int) (*Face, common.Point) {
	for _, otherFace := range m.faces {
		if otherFace.id == face.right {
			for _, p := range m.GetLeftBorder(otherFace) {
				if p.Y == row {
					return &otherFace, p
				}
			}
		}
	}
	panic("no connect left")
}

func (m Map) GetTopBorder(face Face) []common.Point {
	res := []common.Point{}
	for column := face.column * m.FaceWidth(); column < (face.column + 1) * m.FaceWidth(); column++ {
		res = append(res, common.Point{X: column, Y: face.row * m.FaceHeight()})
	}
	return res
}

func (m Map) GetBottomBorder(face Face) []common.Point {
	res := []common.Point{}
	for column := face.column * m.FaceWidth(); column < (face.column + 1) * m.FaceWidth(); column++ {
		res = append(res, common.Point{X: column, Y: ((face.row + 1) * m.FaceHeight()) - 1})
	}
	return res
}

func (m Map) GetLeftBorder(face Face) []common.Point {
	res := []common.Point{}
	for row := face.row * m.FaceHeight(); row < (face.row + 1) * m.FaceHeight(); row++ {
		res = append(res, common.Point{X: face.column * m.FaceWidth(), Y: row})
	}
	return res
}

func (m Map) GetRightBorder(face Face) []common.Point {
	res := []common.Point{}
	for row := face.row * m.FaceHeight(); row < (face.row + 1) * m.FaceHeight(); row++ {
		res = append(res, common.Point{X: ((face.column + 1) * m.FaceWidth()) - 1, Y: row})
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
				otherFace, otherPoint := m.ConnectTop(face, p.X)
				other := board.TileForPoint(otherPoint)				
				t.up = other
				t.upRotation = face.topRotation
				other.down = t
				other.downRotation = otherFace.bottomRotation
			}
		}
		for _, p := range m.GetRightBorder(face) {
			t := board.TileForPoint(p)
			if t.right == nil {
				otherFace, otherPoint := m.ConnectRight(face, p.Y)
				other := board.TileForPoint(otherPoint)				
				t.right = other
				t.rightRotation = face.rightRotation
				other.left = t
				other.leftRotation = otherFace.leftRotation
			}
		} 
		for _, p := range m.GetBottomBorder(face) {
			t := board.TileForPoint(p)
			if t.down == nil {
				otherFace, otherPoint := m.ConnectBottom(face, p.X)
				other := board.TileForPoint(otherPoint)				
				t.down = other
				t.downRotation = face.bottomRotation
				other.up = t
				other.upRotation = otherFace.topRotation
			}
		}
		for _, p := range m.GetLeftBorder(face) {
			t := board.TileForPoint(p)
			if t.left == nil {
				otherFace, otherPoint := m.ConnectLeft(face, p.Y)
				other := board.TileForPoint(otherPoint)				
				t.left = other
				t.leftRotation = face.leftRotation
				other.right = t
				other.rightRotation = otherFace.rightRotation
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
