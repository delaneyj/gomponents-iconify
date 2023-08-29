package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12c0-.341 0-.512-.015-.686a4.042 4.042 0 0 0-.921-2.224a8.285 8.285 0 0 0-.483-.504l-5.167-5.167a8.448 8.448 0 0 0-.504-.483a4.043 4.043 0 0 0-2.224-.92C12.512 2 12.34 2 12 2C7.286 2 4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z"/><path stroke-linecap="round" d="M7 6v3m3-3v3m3-3v3"/></g>`),
		g.Group(children),
	)
}