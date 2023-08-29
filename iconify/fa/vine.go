package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1465 827v198q-101 23-198 23q-65 136-165.5 271T920 1534.5T792 1641q-80 45-162-3q-28-17-60.5-43.5t-85-83.5T382 1382.5t-107.5-184t-105.5-244T77.5 640T7 250h283q26 218 70 398.5t104.5 317T586 1201t140 195q169-169 287-406q-142-72-223-220t-81-333q0-192 104-314.5T1097 0q178 0 273 105.5t95 297.5q0 159-58 286q-7 1-19.5 3t-46 2t-63-6t-62-25.5T1166 611q31-103 31-184q0-87-29-132t-79-45q-53 0-85 49.5T972 440q0 186 105 293.5T1344 841q62 0 121-14z"/>`),
		g.Group(children),
	)
}