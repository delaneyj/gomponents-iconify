package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-digit-zeroa" gradientUnits="userSpaceOnUse" x1="64.01" y1="101.193" x2="64.01" y2="29.735"><stop offset=".485" stop-color="#504f4f"/><stop offset="1" stop-color="#757575"/></linearGradient><path d="M88.2 69.84c0 10.63-1.99 18.65-5.97 24.04c-3.98 5.4-10.03 8.1-18.15 8.1c-7.99 0-14.01-2.63-18.05-7.9c-4.05-5.26-6.12-13.06-6.22-23.38V57.9c0-10.63 1.99-18.6 5.97-23.92c3.98-5.31 10.05-7.97 18.21-7.97c8.09 0 14.12 2.59 18.1 7.77s6.02 12.93 6.12 23.26v12.8h-.01zM75.86 56.03c0-6.94-.95-12.03-2.84-15.26c-1.9-3.23-4.91-4.85-9.04-4.85c-4.03 0-6.99 1.53-8.89 4.6c-1.9 3.06-2.89 7.85-3 14.35v16.76c0 6.91.96 12.05 2.89 15.44s4.96 5.08 9.09 5.08c3.96 0 6.89-1.57 8.79-4.7c1.89-3.13 2.89-8.03 3-14.7V56.03z" fill="url(#ssvg-id-digit-zeroa)"/>`),
		g.Group(children),
	)
}