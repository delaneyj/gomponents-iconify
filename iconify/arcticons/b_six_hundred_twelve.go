package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BSixHundredTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.754 24c1.423 0 2.577 1.137 2.577 2.539s-1.154 2.539-2.577 2.539H9.5V18.923h4.254c1.423 0 2.577 1.137 2.577 2.539S15.177 24 13.754 24Zm-.004.004H9.5"/><ellipse cx="21.778" cy="25.714" rx="3.416" ry="3.364"/><path d="M24.891 20.162c-.569-.734-1.436-1.238-2.873-1.238h-.24c-1.887 0-3.416 1.506-3.416 3.363v3.428m8.508-5.407l2.766-1.381v10.15m2.034-6.791c0-2.076 1.909-3.713 4.092-3.3c1.433.272 2.576 1.492 2.72 2.921c.105 1.064-.236 2.113-.982 2.758c-1.382 1.194-5.83 4.413-5.83 4.413h6.83"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}