package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marvelapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 29.064s.182-7.612 2.434-8.277c2.268-.67.752 4.474 1.072 6.392"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.1 24.403c.1-2.422 1.889-6.01 3.384-4.988c1.073.733-1.212 5.269-.599 7.037m5.834-5.109s-.4-2.424-2.476-.378s-1.587 4.632-.45 4.603c1.857-.049 3.423-5.705 3.423-5.705s-2.15 5.889-.705 5.216s3.163-3.984 3.493-5.741c0 0-1.294 4.462-1.169 5.653c0 0 .978-5.507 3.287-5.795c0 0 .421-.441-.126 1.292s.1 2.15 2.895-1.25c0 0-2.296 4.787-.812 4.923s3.925-3.932 3.83-4.938s-1.81.863-.158 2.19s4.988-.621 4.552-1.773s-2.93-.853-3.533 2.7s2.804 1.003 5.065-.61s4.253-3.366 3.505-4.446c-.755-1.088-3.08 1.15-3.885 5.326s2.703.15 2.703.15m-3.315 3.734s-18.731-.673-25.554 4.495"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}