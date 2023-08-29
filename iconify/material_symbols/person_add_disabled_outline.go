package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonAddDisabledOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L17 19.8v.2H1v-2.8q0-.85.438-1.563T2.6 14.55q1.55-.775 3.15-1.163T9 13q.325 0 .638.013t.612.037L9.2 12H9q-1.65 0-2.825-1.175T5 8v-.2L1.4 4.2l1.425-1.425l18.4 18.4L19.8 22.6ZM18 14v-3h-3V9h3V6h2v3h3v2h-3v3h-2Zm-5.4-4.3L11 8.1V8q0-.825-.587-1.412T9 6h-.1L7.3 4.4q.375-.2.813-.3T9 4q1.65 0 2.825 1.175T13 8q0 .45-.1.888t-.3.812ZM3 18h12l-2.15-2.35q-.95-.325-1.925-.488T9 15q-1.4 0-2.775.338T3.5 16.35q-.225.125-.363.35T3 17.2v.8Zm8-9.9ZM9 18Z"/>`),
		g.Group(children),
	)
}