package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherFogFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M32.25 38.5a1.75 1.75 0 0 1 .144 3.494l-.143.006h-16.5a1.75 1.75 0 0 1-.144-3.494l.144-.006h16.5Zm6-5.995a1.75 1.75 0 0 1 .144 3.494l-.143.006H9.75a1.75 1.75 0 0 1-.143-3.494l.143-.006h28.5ZM24 6.01c6.337 0 9.932 4.194 10.455 9.26h.16c4.078 0 7.384 3.297 7.384 7.365S38.692 30 34.614 30h-21.23C9.306 30 6 26.703 6 22.635s3.306-7.365 7.384-7.365h.16C14.07 10.171 17.662 6.01 24 6.01Z"/>`),
		g.Group(children),
	)
}