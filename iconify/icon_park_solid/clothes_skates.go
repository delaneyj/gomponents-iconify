package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesSkates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClothesSkates0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 4H8c-1.886 0-2.828 0-3.414.586C4 5.172 4 6.114 4 8v32c0 1.886 0 2.828.586 3.414C5.172 44 6.114 44 8 44h4l4-4l4 4h20c1.886 0 2.828 0 3.414-.586C44 42.828 44 41.886 44 40v-7.525c0-4.992 0-7.488-1.48-9.183c-1.48-1.696-3.953-2.033-8.899-2.707l-1.242-.17c-4.946-.674-7.42-1.011-8.9-2.707C22 16.013 22 13.518 22 8.525V8c0-1.886 0-2.828-.586-3.414C20.828 4 19.886 4 18 4Zm26 30H4m0-22.5h18"/><circle cx="15" cy="23" r="3" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClothesSkates0)"/>`),
		g.Group(children),
	)
}