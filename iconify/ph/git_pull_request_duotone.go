package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequestDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M96 64a24 24 0 1 1-24-24a24 24 0 0 1 24 24Zm104 104a24 24 0 1 0 24 24a24 24 0 0 0-24-24Z" opacity=".2"/><path d="M104 64a32 32 0 1 0-40 31v66a32 32 0 1 0 16 0V95a32.06 32.06 0 0 0 24-31Zm-48 0a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm32 128a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm120-31v-37.12a55.67 55.67 0 0 0-16.4-39.6L163.31 56H192a8 8 0 0 0 0-16h-48a8 8 0 0 0-8 8v48a8 8 0 0 0 16 0V67.31l28.28 28.29A39.71 39.71 0 0 1 192 123.88V161a32 32 0 1 0 16 0Zm-8 47a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/></g>`),
		g.Group(children),
	)
}