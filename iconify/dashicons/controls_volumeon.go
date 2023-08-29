package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlsVolumeon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 7h4l5-4v14l-5-4H2V7zm12.69-2.46C14.82 4.59 18 5.92 18 10s-3.18 5.41-3.31 5.46a.489.489 0 0 1-.65-.27c-.11-.26.02-.55.27-.65c.11-.05 2.69-1.15 2.69-4.54c0-3.41-2.66-4.53-2.69-4.54a.493.493 0 0 1-.27-.65c.1-.25.39-.38.65-.27zM16 10c0 2.57-2.23 3.43-2.32 3.47c-.06.02-.12.03-.18.03c-.2 0-.39-.12-.47-.32c-.1-.26.04-.55.29-.65c.07-.02 1.68-.67 1.68-2.53s-1.61-2.51-1.68-2.53a.52.52 0 0 1-.29-.65c.1-.25.39-.39.65-.29c.09.04 2.32.9 2.32 3.47z"/>`),
		g.Group(children),
	)
}