package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Biking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M400 96a48 48 0 1 0-48-48a48 48 0 0 0 48 48zm-4 121a31.9 31.9 0 0 0 20 7h64a32 32 0 0 0 0-64h-52.78L356 103a31.94 31.94 0 0 0-40.81.68l-112 96a32 32 0 0 0 3.08 50.92L288 305.12V416a32 32 0 0 0 64 0V288a32 32 0 0 0-14.25-26.62l-41.36-27.57l58.25-49.92zm116 39a128 128 0 1 0 128 128a128 128 0 0 0-128-128zm0 192a64 64 0 1 1 64-64a64 64 0 0 1-64 64zM128 256a128 128 0 1 0 128 128a128 128 0 0 0-128-128zm0 192a64 64 0 1 1 64-64a64 64 0 0 1-64 64z"/>`),
		g.Group(children),
	)
}