package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.36 11.33l-1.1-.25l.86 1.22v3.79l2.93-2.44V9.5l-1.35.48l-1.34 1.35zm-6.84 0L7.17 9.98L5.83 9.5v4.15l2.93 2.44V12.3l.86-1.22l-1.1.25z"/><path fill="currentColor" d="M13.16 11.45h-2.44l-.61-.37l-.98 1.47v5.5l.73 1.09l.86.86h2.44l.86-.86l.73-1.09v-5.5l-.98-1.47l-.61.37zm1.96 6.96l1.58-1.59v-1.58l-1.58 1.34v1.83zm-7.95-1.59l1.59 1.59v-1.83l-1.59-1.34v1.58zm3.91-5.86h.62V4h-1.35l-.97 2.31l-4.41-.36l.74 3.06l5.25 1.95h.12zm3.42-4.65L13.53 4h-1.35v7H13l5.25-2L19 6Z"/><path fill="currentColor" d="M17.32 5.71L15.6 4h-1.71l.86 1.95l2.57-.24zM9.98 4H8.27L6.56 5.71l2.57.24L9.98 4z"/>`),
		g.Group(children),
	)
}