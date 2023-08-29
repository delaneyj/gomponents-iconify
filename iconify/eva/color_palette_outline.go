package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorPaletteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaColorPaletteOutline0"><g id="evaColorPaletteOutline1"><g id="evaColorPaletteOutline2" fill="currentColor"><path d="M19.54 5.08A10.61 10.61 0 0 0 11.91 2a10 10 0 0 0-.05 20a2.58 2.58 0 0 0 2.53-1.89a2.52 2.52 0 0 0-.57-2.28a.5.5 0 0 1 .37-.83h1.65A6.15 6.15 0 0 0 22 11.33a8.48 8.48 0 0 0-2.46-6.25ZM15.88 15h-1.65a2.49 2.49 0 0 0-1.87 4.15a.49.49 0 0 1 .12.49c-.05.21-.28.34-.59.36a8 8 0 0 1-7.82-9.11A8.1 8.1 0 0 1 11.92 4H12a8.47 8.47 0 0 1 6.1 2.48a6.5 6.5 0 0 1 1.9 4.77A4.17 4.17 0 0 1 15.88 15Z"/><circle cx="12" cy="6.5" r="1.5"/><path d="M15.25 7.2a1.5 1.5 0 1 0 2.05.55a1.5 1.5 0 0 0-2.05-.55Zm-6.5 0a1.5 1.5 0 1 0 .55 2.05a1.5 1.5 0 0 0-.55-2.05Zm-2.59 4.06a1.5 1.5 0 1 0 2.08.4a1.49 1.49 0 0 0-2.08-.4Z"/></g></g></g>`),
		g.Group(children),
	)
}