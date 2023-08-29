package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M7 0C4.795 0 3 1.795 3 4v18c0 2.205 1.795 4 4 4h12c2.205 0 4-1.795 4-4V8c0-1.062-.973-2.07-2.719-3.781c-.244-.24-.504-.502-.75-.75c-.248-.246-.51-.475-.75-.719C17.07 1.003 16.063 0 15 0H7zm0 2h7.281c.721.184.719 1.05.719 1.938V7c0 .551.449 1 1 1h3c.998 0 2 .005 2 1v13a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2zm3 7a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm5.969 3.156c-1.187 0-1.607 4.906-3.25 4.906c-1.217 0-1.69-2-2.719-2S7.125 19.25 7.125 19.25c-.38.422.186.75.625.75h10.563c.357 0 .81-.453.53-1.063c0 0-1.687-6.78-2.874-6.78z"/>`),
		g.Group(children),
	)
}