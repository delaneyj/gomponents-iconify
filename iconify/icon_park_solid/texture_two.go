package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextureTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTextureTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M40 4H8a4 4 0 0 0-4 4v32a4 4 0 0 0 4 4h32a4 4 0 0 0 4-4V8a4 4 0 0 0-4-4Z"/><path stroke="#000" d="M12 12v8m16 8v8m-8-24v8m8-8h8M12 28h8m8-8h8M12 36h8m16-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTextureTwo0)"/>`),
		g.Group(children),
	)
}