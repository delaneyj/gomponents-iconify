package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gtn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2h0ZM23.608 31.261c-5.72 3.304-10.663 4.151-12.194 2.09c-.395-.531-.546-1.24-.445-2.093m13.423-14.52c5.72-3.303 10.663-4.15 12.194-2.09c.395.532.546 1.24.445 2.094M31.567 29.23V18.764L38.5 29.229V18.764M24.098 29.229V18.764m-3.532 0H27.5m-11.066 3.54c0-1.963-1.7-3.664-3.663-3.533c-1.831.131-3.27 1.832-3.27 3.663v3.27a3.517 3.517 0 0 0 3.532 3.532h0a3.517 3.517 0 0 0 3.532-3.532h-3.532"/>`),
		g.Group(children),
	)
}