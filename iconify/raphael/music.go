package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.75 8.042v15.903c-.676-.45-1.635-.736-2.707-.736c-2.048 0-3.708 1.024-3.708 2.29s1.66 2.292 3.708 2.292s3.708-1.026 3.708-2.29V13.785l10.916-3.24v9.565c-.678-.45-1.635-.735-2.708-.735c-2.048 0-3.708 1.025-3.708 2.292c0 1.265 1.66 2.29 3.708 2.29s3.708-1.025 3.708-2.29V4.207L12.75 8.042z"/>`),
		g.Group(children),
	)
}