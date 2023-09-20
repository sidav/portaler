package portaler

func (sc *Scene) InitTesting2() {
	// sector 1
	s := &sector{floorHeight: -0.5, ceilingHeight: 0.5, lightLevel: 255}
	s.addLinedef(-10, -5, -10, 5)
	s.appendLinedef(-5, 5)
	s.appendLinedef(-5, -1)
	s.appendLinedef(-5, -3)
	s.appendLinedef(-5, -5)
	s.appendLinedef(-10, -5)
	sc.addSector(s)

	s = &sector{floorHeight: 0, ceilingHeight: 1.5, lightLevel: 160}
	s.addLinedef(-5, -5, -5, -3)
	s.appendLinedef(-5, -1)
	s.appendLinedef(-5, 5)
	s.appendLinedef(0, 5)
	s.appendLinedef(5, 5)
	s.appendLinedef(5, -5)
	s.appendLinedef(-5, -5)
	sc.addSector(s)
	sc.portalizeLinedefWithCoordinates(-5, -1, -5, -3)

	s = &sector{floorHeight: 0, ceilingHeight: 0.75, lightLevel: 128}
	s.addLinedef(-5, 5, -5, 7)
	s.appendLinedef(-5, 9)
	s.appendLinedef(-5, 10)
	s.appendLinedef(0, 10)
	s.appendLinedef(0, 8)
	s.appendLinedef(0, 6)
	s.appendLinedef(0, 5)
	s.appendLinedef(-5, 5)
	sc.addSector(s)
	sc.portalizeLinedefWithCoordinates(-5, 5, 0, 5)

	s = &sector{floorHeight: 0.25, ceilingHeight: 1.25, lightLevel: 96}
	s.addLinedef(-10, 5, -10, 10)
	s.appendLinedef(-5, 10)
	s.appendLinedef(-5, 9)
	s.appendLinedef(-5, 7)
	s.appendLinedef(-5, 5)
	s.appendLinedef(-10, 5)
	sc.addSector(s)
	sc.portalizeLinedefWithCoordinates(-5, 7, -5, 9)

	s = &sector{floorHeight: 0, ceilingHeight: 5.5, lightLevel: 64}
	s.addLinedef(0, 5, 0, 6)
	s.appendLinedef(0, 8)
	s.appendLinedef(0, 10)
	s.appendLinedef(5, 10)
	s.appendLinedef(5, 5)
	s.appendLinedef(0, 5)
	sc.addSector(s)
	sc.portalizeLinedefWithCoordinates(0, 6, 0, 8)
	sc.scale(0.5)
}

func (sc *Scene) InitTesting1() {
	// sector 1
	s := &sector{floorHeight: 0, ceilingHeight: 3, lightLevel: 255}
	s.addLinedef(-4, 0, -1, -5)
	s.appendLinedef(4, -4)
	s.appendLinedef(4, 2)
	s.appendLinedef(0, 5)
	s.appendLinedef(-4, 0)
	sc.addSector(s)

	s = &sector{floorHeight: -2, ceilingHeight: 2, lightLevel: 128}
	s.addLinedef(4, -4, 10, -4)
	s.appendLinedef(10, 2)
	s.appendLinedef(4, 2)
	s.appendLinedef(4, -4)
	sc.addSector(s)
	sc.portalizeLinedefWithCoordinates(4, -4, 4, 2)

	s = &sector{floorHeight: 1, ceilingHeight: 4, lightLevel: 64}
	s.addLinedef(4, 2, 10, 2)
	s.appendLinedef(10, 8)
	s.appendLinedef(4, 8)
	s.appendLinedef(4, 2)
	sc.addSector(s)
	sc.portalizeLinedefWithCoordinates(4, 2, 10, 2)
}
