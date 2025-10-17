package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/area"
)

func (gd *GameReader) TerrorZones() (areas []area.ID) {
	tzPtr := uintptr(gd.ReadUInt(gd.moduleBaseAddressPtr+gd.offset.TZ, Uint64))

	for i := 0; i < 8; i++ {
		tzArea := gd.ReadUInt(tzPtr+uintptr(i*Uint32), Uint32)
		if tzArea != 0 {
			areas = append(areas, area.ID(tzArea))
		}
	}

	return
}
