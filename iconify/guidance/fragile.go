package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fragile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m12 18.939l.002.086v4.451M12 18.94c-.055-.985-1.055-2.187-2.998-2.969c-2.786-1.112-4.506-3.772-4.504-6.876c0-.6.073-1.197.217-1.78L6.502.524h11l1.783 6.79a7.42 7.42 0 0 1 .217 1.78c.002 3.104-1.718 5.764-4.504 6.876c-1.943.782-2.943 1.984-2.998 2.969Zm.002 4.537h6m-6 0h-6m7.5-13.45l.232-.233a6.035 6.035 0 0 0 1.768-4.268h-3.5v-.25l.232-.23a6.04 6.04 0 0 0 1.768-4.27v-.25"/>`),
		g.Group(children),
	)
}