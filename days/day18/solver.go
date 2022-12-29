package day18

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	cubes := parseCubes(input)
	neighbors := countSideNeighbors(cubes)
	return strconv.Itoa(len(cubes)*6 - neighbors*2)
}

func countSideNeighbors(cubes []cube) int {
	count := 0
	for i, cube := range cubes {
		for j := i + 1; j < len(cubes); j++ {
			if cube.isSideNeighbor(cubes[j]) {
				count++
			}
		}
	}
	return count
}

func parseCubes(input string) []cube {
	lines := strings.Split(input, "\n")
	cubes := make([]cube, len(lines))
	for i, line := range lines {
		cubes[i] = parseCube(line)
	}
	return cubes
}

type cube struct {
	x, y, z int
}

func parseCube(line string) cube {
	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return cube{x, y, z}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c *cube) isSideNeighbor(other cube) bool {
	if c.x == other.x && c.y == other.y {
		return abs(c.z-other.z) == 1
	}
	if c.x == other.x && c.z == other.z {
		return abs(c.y-other.y) == 1
	}
	if c.y == other.y && c.z == other.z {
		return abs(c.x-other.x) == 1
	}
	return false
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	fullCube := fullCube(parseCubes(input))
	d := createDroplet(fullCube)
	d.fillCubes(parseCubes(input))
	d.setExterior()
	d.fill()
	cubes := d.toCubes()
	neighbors := countSideNeighbors(cubes)
	return strconv.Itoa(len(cubes)*6 - neighbors*2)
}

func fullCube(cubes []cube) cube {
	maxX, maxY, maxZ := 0, 0, 0
	for _, cube := range cubes {
		if cube.x > maxX {
			maxX = cube.x
		}
		if cube.y > maxY {
			maxY = cube.y
		}
		if cube.z > maxZ {
			maxZ = cube.z
		}
	}
	return cube{maxX, maxY, maxZ}
}

type space int8

const (
	UNCHECKED space = iota
	CUBE
	EXTERIOR
)

type droplet [][][]space

func createDroplet(c cube) droplet {
	full := make(droplet, c.x+1)
	for i := range full {
		full[i] = make([][]space, c.y+1)
		for j := range full[i] {
			full[i][j] = make([]space, c.z+1)
		}
	}
	return full
}

func (d *droplet) fillCubes(cubes []cube) {
	for _, cube := range cubes {
		(*d)[cube.x][cube.y][cube.z] = CUBE
	}
}

func (d *droplet) setExterior() {
	xMax := len(*d) - 1
	for y := range (*d)[0] {
		for z := range (*d)[0][y] {
			if (*d)[0][y][z] == UNCHECKED {
				(*d)[0][y][z] = EXTERIOR
			}
			if (*d)[xMax][y][z] == UNCHECKED {
				(*d)[xMax][y][z] = EXTERIOR
			}
		}
	}
	yMax := len((*d)[0]) - 1
	for x := range *d {
		for z := range (*d)[x][0] {
			if (*d)[x][0][z] == UNCHECKED {
				(*d)[x][0][z] = EXTERIOR
			}
			if (*d)[x][yMax][z] == UNCHECKED {
				(*d)[x][yMax][z] = EXTERIOR
			}
		}
	}
	zMax := len((*d)[0][0]) - 1
	for x := range *d {
		for y := range (*d)[x] {
			if (*d)[x][y][0] == UNCHECKED {
				(*d)[x][y][0] = EXTERIOR
			}
			if (*d)[x][y][zMax] == UNCHECKED {
				(*d)[x][y][zMax] = EXTERIOR
			}
		}
	}
}

func (d *droplet) countUnchecked() int {
	count := 0
	for x := range *d {
		for y := range (*d)[x] {
			for z := range (*d)[x][y] {
				if (*d)[x][y][z] == UNCHECKED {
					count++
				}
			}
		}
	}
	return count
}

func (d *droplet) fill() {
	oldUnchecked := 0
	unchecked := d.countUnchecked()
	for unchecked != oldUnchecked {
		oldUnchecked = unchecked
		for xi := 1; xi < len(*d)-1; xi++ {
			for yi := 1; yi < len((*d)[xi])-1; yi++ {
				for zi := 1; zi < len((*d)[xi][yi])-1; zi++ {
					if (*d)[xi][yi][zi] == UNCHECKED && d.hasExteriorNeighbor(xi, yi, zi) {
						(*d)[xi][yi][zi] = EXTERIOR
					}
				}
			}
		}
		unchecked = d.countUnchecked()
	}
}

func (d *droplet) hasExteriorNeighbor(x, y, z int) bool {
	if (*d)[x-1][y][z] == EXTERIOR {
		return true
	}
	if (*d)[x+1][y][z] == EXTERIOR {
		return true
	}
	if (*d)[x][y-1][z] == EXTERIOR {
		return true
	}
	if (*d)[x][y+1][z] == EXTERIOR {
		return true
	}
	if (*d)[x][y][z-1] == EXTERIOR {
		return true
	}
	if (*d)[x][y][z+1] == EXTERIOR {
		return true
	}
	return false
}

func (d *droplet) toCubes() []cube {
	cubes := make([]cube, 0, len(*d)*len((*d)[0])*len((*d)[0][0]))
	for x := range *d {
		for y := range (*d)[x] {
			for z := range (*d)[x][y] {
				if (*d)[x][y][z] != EXTERIOR {
					cubes = append(cubes, cube{x, y, z})
				}
			}
		}
	}
	return cubes
}
