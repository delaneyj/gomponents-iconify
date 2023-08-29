package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vlc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.913 18.8a10.64 10.64 0 0 0 5.768-1.97l2.82 8.44a14.932 14.932 0 0 1-8.647 2.99a14.932 14.932 0 0 1-8.646-2.99l2.82-8.44a10.64 10.64 0 0 0 5.769 1.97Zm-.058-14.3c-1.413 0-2.084 1.377-2.397 2.332l-.983 2.866a5.767 5.767 0 0 0 3.38.961a5.767 5.767 0 0 0 3.38-.96l-.983-2.867c-.312-.955-.984-2.332-2.397-2.332Zm0 32.063a16.552 16.552 0 0 1-11.208-4.033l1.122-3.272h-2.818c-1.093 0-1.538.363-1.814 1.22l-3.24 10.949C5.574 42.69 5.707 43.5 7.16 43.5h33.392c1.453 0 1.585-.81 1.261-2.073l-3.24-10.948c-.275-.858-.72-1.221-1.813-1.221h-2.818l1.122 3.272a16.552 16.552 0 0 1-11.208 4.033Z"/>`),
		g.Group(children),
	)
}