package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Esound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.977 30.723c2.568 0 2.857-1.553 2.857-2.441c0-2.387-2.233-4.308-5.122-4.308a5.122 5.122 0 0 0-5.121 5.122c0 1.744 1 2.786 1 5.493s-1.389 8.02-8.764 8.02"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M27.977 30.723c-1.419 0-2.569 1.15-2.569 2.568s.972 2.568 2.568 2.568c2.604 0 7.88-4.36 12.158-7.874c1.161-.954 2.069-2.261 2.069-3.877a4.639 4.639 0 0 0-1.26-3.17s-.147-.157-.193-.2c-8.273-8.402-21.834-13.942-26.741-14.96a4.643 4.643 0 0 0-.68-.12a5.32 5.32 0 0 0-.502-.05a4.668 4.668 0 0 0-4.668 4.668c0 .738.186 1.427.491 2.047l-.01.017c1.834 3.307 5.263 10.053 7.273 16.683l.001-.007"/><path d="M12.827 33.272c-3.17.406-4.668 2.09-4.668 4.668a4.668 4.668 0 1 0 9.336 0h0c0-2.7-.635-5.795-1.581-8.917c-.01.076.226 3.825-3.087 4.25Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.977 35.86c2.425 0 3.641-3.614 2.857-7.935"/>`),
		g.Group(children),
	)
}