package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M44 41V19H38V41H44Z"/><path d="M38 19.0002C30.8948 12.4304 26.7757 8.66359 25.6428 7.69983C23.9433 6.25419 22.0226 6.86123 22.0226 10.479C22.0226 14.0968 27.2864 16.2443 27.2864 19.0002C27.2898 19.0166 20.529 19.0177 7.00404 19.0035C5.3467 19.0017 4.00175 20.3438 4 22.0012C4 22.0022 4 22.0033 4 22.0043C4 23.6635 5.34501 25.0085 7.00417 25.0085H14.0165C15.2234 32.9771 15.8893 37.3101 16.0144 38.0075C16.2019 39.0536 17.199 41.0002 20.068 41.0002C21.9807 41.0002 27.9581 41.0002 38 41.0002V19.0002Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}