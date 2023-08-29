package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodepenLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234 92.55s-.05 0-.09-.07l-104-56a4 4 0 0 0-3.8 0l-104 56l-.11.08A4 4 0 0 0 20 96v64a4 4 0 0 0 2 3.45a.71.71 0 0 0 .09.07l104 56a4 4 0 0 0 3.8 0l104-56a.27.27 0 0 0 .08-.07a4 4 0 0 0 2-3.45V96a4 4 0 0 0-1.97-3.45Zm-6 60.75L181 128l47-25.3Zm-55.43-29.84L132 101.61V46.7L223.56 96Zm-44.57 24L91.87 128L128 108.54L164.13 128ZM124 46.7v54.91l-40.57 21.85L32.44 96Zm-96 56L75 128l-47 25.3Zm55.43 29.84L124 154.39v54.91L32.44 160ZM132 209.3v-54.91l40.57-21.85l51 27.46Z"/>`),
		g.Group(children),
	)
}