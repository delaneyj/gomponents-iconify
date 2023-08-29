package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TakeoutBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M51 27h4v2h-4zm-34 0h4v2h-4zm3 22h32v2H20z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17.712 29.958l2.85 24.542h30.875l2.85-24.542"/><path d="M54.269 29.956L51.437 54.5H20.562L17 23.625m38 0l-.73 6.331M15 22L8 37m49-15l7 15M20 22h32m-32 0l2 15m30-15l-2 15m-28 0h28M21 51h30m-31-3h31M18 26h2m-2 3h2m32-3h2m-2 3h2m-12-7l12.709-11.276a.961.961 0 0 0-1.359-1.36zm-4 0L49.32 6.549a.955.955 0 0 0-.29-1.359a.997.997 0 0 0-1.36.291z"/></g>`),
		g.Group(children),
	)
}