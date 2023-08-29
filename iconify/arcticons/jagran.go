package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jagran(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.303 26.19c1.21 2.415 1.95 4.857 2.1 7.705H4.597c.227-10.968 9.189-19.79 20.21-19.79c7.916 0 14.769 4.55 18.087 11.176l-9.62 4.792"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.642 33.807c.204-9.862 8.262-17.68 18.172-17.68c4.554 0 8.716 1.675 11.905 4.441l-5.766 6.565"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.076 33.728a12.904 12.904 0 0 1-.018-.683c0-7.035 5.703-12.738 12.738-12.738c1.327 0 2.606.203 3.808.579l-1.167 3.992"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.754 33.73a9.05 9.05 0 0 1 18.1 0"/>`),
		g.Group(children),
	)
}