package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesGloves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTClothesGloves0"><g fill="none"><path fill="#555" d="M27 4H15c-3.771 0-5.657 0-6.828 1.172C7 6.343 7 8.229 7 12v32h28v-7s7 0 7-6v-8c0-6-7-6-7-6v-5c0-3.771 0-5.657-1.172-6.828C32.657 4 30.771 4 27 4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 27V17m0 0v-5c0-3.771 0-5.657-1.172-6.828C32.657 4 30.771 4 27 4H15c-3.771 0-5.657 0-6.828 1.172C7 6.343 7 8.229 7 12v32h28v-7s7 0 7-6v-8c0-6-7-6-7-6Zm-21 5V4m7 18V4m7 18V4M12 4h18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTClothesGloves0)"/>`),
		g.Group(children),
	)
}