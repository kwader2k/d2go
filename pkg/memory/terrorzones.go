package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/area"
)

func (gd *GameReader) TerrorZones() (areas []area.ID) {
	structPtr := gd.moduleBaseAddressPtr + gd.offset.TZ
	zonesPtr := uintptr(gd.ReadUInt(structPtr, Uint64))
	actualActiveZoneCount := int(gd.ReadUInt(structPtr+0x8, Uint8))

	for i := 0; i < actualActiveZoneCount; i++ {
		tzArea := gd.ReadUInt(zonesPtr+uintptr(i*Uint32), Uint32)
		if tzArea != 0 {
			areas = append(areas, area.ID(tzArea))
		}
	}

	return
}
