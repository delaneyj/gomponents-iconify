package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Direct(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M74.531 24.638V41.34H47.826v17.318h26.705v16.703L100 50Zm-56.416 7.247C8.153 31.885 0 40.038 0 50c0 9.962 8.153 18.115 18.115 18.115c9.714 0 17.651-7.768 18.043-17.39a3.623 3.623 0 0 0 0-.004a3.623 3.623 0 0 0 .074-.721a3.623 3.623 0 0 0-.074-.72c-.39-9.625-8.327-17.395-18.043-17.395Zm0 7.246A10.816 10.816 0 0 1 28.986 50c0 6.046-4.825 10.87-10.87 10.87A10.814 10.814 0 0 1 7.245 50c0-6.046 4.823-10.87 10.87-10.87z" color="currentColor"/>`),
		g.Group(children),
	)
}