package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TomlSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 .774C0 .347.347 0 .774 0h3.613a.774.774 0 0 1 0 1.548H1.548v12.904h2.84a.774.774 0 1 1 0 1.548H.773A.774.774 0 0 1 0 15.226V.774Zm10.839 0c0-.427.346-.774.774-.774h3.613c.427 0 .774.347.774.774v14.452a.774.774 0 0 1-.774.774h-3.613a.774.774 0 1 1 0-1.548h2.839V1.548h-2.84a.774.774 0 0 1-.773-.774Zm-6.71 4.13c0-.428.347-.775.774-.775h6.194a.774.774 0 0 1 0 1.548H8.774v7.484a.774.774 0 1 1-1.548 0V5.677H4.903a.774.774 0 0 1-.774-.774Z"/>`),
		g.Group(children),
	)
}