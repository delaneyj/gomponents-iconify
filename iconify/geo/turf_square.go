package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M37.043 74.903a2 2 0 0 1-2-2v-9.161a2 2 0 0 1 4 0v9.161a2 2 0 0 1-2 2zm0-18.322a2 2 0 0 1-2-2v-9.162a2 2 0 0 1 4 0v9.162a2 2 0 0 1-2 2zm0-18.323a2 2 0 0 1-2-2v-9.161a2 2 0 0 1 4 0v9.161a2 2 0 0 1-2 2zm25.914 36.645a2 2 0 0 1-2-2v-9.161a2 2 0 0 1 4 0v9.161a2 2 0 0 1-2 2zm0-18.323a2 2 0 0 1-2-2v-9.161a2 2 0 0 1 4 0v9.161a2 2 0 0 1-2 2zm0-18.322a2 2 0 0 1-2-2v-9.161a2 2 0 0 1 4 0v9.161a2 2 0 0 1-2 2z"/><path fill="none" d="M63 19.936a2 2 0 0 1-2-2V15H39v2.936a2 2 0 0 1-4 0V15H15v70h20v-2.936a2 2 0 0 1 4 0V85h22v-2.936a2 2 0 0 1 4 0V85h20V15H65v2.936a2 2 0 0 1-2 2z"/><path fill="currentColor" d="M87.064 11H12.936C11.831 11 11 11.831 11 12.936v74.129c0 1.104.831 1.935 1.936 1.935h74.129C88.169 89 89 88.169 89 87.064V12.936C89 11.831 88.169 11 87.064 11zM85 85H65v-2.936a2 2 0 0 0-4 0V85H39v-2.936a2 2 0 0 0-4 0V85H15V15h20v2.936a2 2 0 0 0 4 0V15h22v2.936a2 2 0 0 0 4 0V15h20v70z"/>`),
		g.Group(children),
	)
}