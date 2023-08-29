package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creativeapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.746 45c-3.15-.038-16.517-7.973-18.06-10.72s-1.353-18.291.254-21S21.104 2.962 24.255 3s16.517 7.973 18.06 10.72s1.353 18.291-.255 21S26.896 45.038 23.747 45Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.96 34.99l-9.339-5.558V18.315l9.339-5.558l9.339 5.558V29.43Zm9.339-16.675l10.069 5.936m-28.747-5.936l.025-11.276M23.96 34.99l-9.86 5.658"/>`),
		g.Group(children),
	)
}