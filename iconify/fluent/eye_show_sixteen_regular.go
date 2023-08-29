package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeShowSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M2.984 8.625v.003a.5.5 0 0 1-.612.355c-.431-.114-.355-.611-.355-.611l.017-.062s.026-.084.047-.145a6.7 6.7 0 0 1 1.117-1.982C4.097 5.089 5.606 4 8 4c2.395 0 3.904 1.089 4.801 2.183a6.7 6.7 0 0 1 1.117 1.982a3.916 3.916 0 0 1 .06.187l.003.013l.002.004v.002a.5.5 0 0 1-.966.258l-.001-.004l-.008-.025a4.9 4.9 0 0 0-.2-.52a5.703 5.703 0 0 0-.78-1.263C11.285 5.912 10.044 5 8 5c-2.044 0-3.286.912-4.028 1.817a5.701 5.701 0 0 0-.945 1.674a3.018 3.018 0 0 0-.035.109l-.008.025z" fill="currentColor"/><path d="M8 7a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5zM6.5 9.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}