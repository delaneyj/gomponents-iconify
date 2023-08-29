package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vrbankingclassic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.484 30.337H43.5v-19.5h-39v19.5h14.016l-1.828-8.53h-4.875l-1.22-4.876h10.36L24 30.337l2.072-9.75c.703-3.31 2.232-3.656 5.24-3.656h6.094l-1.584 7.313a2.058 2.058 0 0 1-1.95 1.584c-.689-.064-1.721-.656-1.585-1.34l.366-1.829a.719.719 0 0 0-.731-.853a.943.943 0 0 0-.853.731Zm-24.984 0v6.825h15.478l-1.462-6.824m10.968-.001l-1.462 6.825H43.5v-6.824"/>`),
		g.Group(children),
	)
}