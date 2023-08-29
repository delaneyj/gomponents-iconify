package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jyskebank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.984 14.081c1.37-3.299 5.647-7.536 8.951-9.06l-.03 17.9ZM17.48 3.516c3.296-1.362 9.31-1.326 12.72-.062L17.553 16.09ZM33.927 4.91c3.291 1.371 7.519 5.66 9.039 8.97l-17.859-.03Zm10.559 12.663c1.36 3.304 1.324 9.332.062 12.75L31.941 17.645Zm-1.469 16.345c-1.37 3.3-5.647 7.536-8.951 9.06l.03-17.9ZM30.52 44.484c-3.296 1.362-9.31 1.326-12.72.062L30.448 31.91ZM14.073 43.09c-3.291-1.372-7.518-5.66-9.039-8.971l17.859.03ZM3.514 30.426c-1.36-3.303-1.324-9.331-.062-12.75l12.607 12.679Z"/>`),
		g.Group(children),
	)
}