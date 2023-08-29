package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ambulance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(2 2)"><circle cx="1.433" cy="9.433" r="1.433"/><ellipse cx="9.451" cy="9.474" rx="1.451" ry="1.474"/><path d="M8.953 2.047c0-1.122.1-2.031-.944-2.031c-1.045 0-.993.909-.993 2.031h1.937z"/></g><path d="M15.338 7.735S12.729 4 11.994 4H1.392C.659 4 .064 4.599.064 5.336v5.315c0 .736.328 1.336.328 1.336h.553a2.656 2.656 0 0 1-.043-.431a2.6 2.6 0 1 1 5.2 0c0 .147-.02.29-.043.431h2.887a2.656 2.656 0 0 1-.043-.431a2.6 2.6 0 1 1 5.2 0c0 .147-.02.29-.043.431h.594c.734 0 1.329-.6 1.329-1.336v-1.33c-.002-.834-.645-1.586-.645-1.586zm-7.291-.688H7.043v1.005H5.958V7.047H4.953V5.963h1.005V4.958h1.085v1.005h1.004v1.084zm1.871.995V4.904h1.367c.636 0 1.375 1.031 1.375 1.031l1.464 2.106H9.918v.001z"/></g>`),
		g.Group(children),
	)
}