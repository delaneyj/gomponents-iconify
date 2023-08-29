package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#DA1157"/><path fill="#FFF" d="M20.985 19.168a3.026 3.026 0 1 0 0-6.052a3.026 3.026 0 0 0 0 6.052zm0-10.005a6.98 6.98 0 1 1-5.003 11.843a6.955 6.955 0 0 0 1.977-4.864a6.954 6.954 0 0 0-1.977-4.864a6.957 6.957 0 0 1 5.003-2.115zM10.98 19.168a3.026 3.026 0 1 0 0-6.052a3.026 3.026 0 0 0 0 6.052zm5.003-7.89a6.955 6.955 0 0 0-1.976 4.864c0 1.892.754 3.607 1.976 4.864a6.98 6.98 0 1 1 0-9.728zm-1.977 4.865c0 1.892.754 3.607 1.977 4.864a6.954 6.954 0 0 0 1.976-4.864a6.954 6.954 0 0 0-1.976-4.864a6.954 6.954 0 0 0-1.977 4.864z"/></g>`),
		g.Group(children),
	)
}