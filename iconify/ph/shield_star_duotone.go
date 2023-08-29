package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldStarDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 56v58.77c0 84.18-71.31 112.07-85.54 116.8a7.54 7.54 0 0 1-4.92 0C111.31 226.86 40 199 40 114.79V56a8 8 0 0 1 8-8h160a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M80.57 117A8 8 0 0 1 91 112.57l29 11.61V96a8 8 0 0 1 16 0v28.18l29-11.61a8 8 0 1 1 6 14.86l-30.31 12.12l17.71 23.65a8 8 0 1 1-12.8 9.6L128 149.33l-17.6 23.47a8 8 0 1 1-12.8-9.6l17.74-23.65L85 127.43A8 8 0 0 1 80.57 117ZM224 56v58.77c0 89.62-75.82 119.34-91 124.39a15.44 15.44 0 0 1-10 0c-15.2-5-91-34.76-91-124.38V56a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Zm-16 0H48v58.79c0 78.5 66.47 104.68 80 109.18c13.66-4.56 80-30.76 80-109.18Z"/></g>`),
		g.Group(children),
	)
}