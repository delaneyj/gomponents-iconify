package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6.879L17.121 3A1.497 1.497 0 0 0 15 3L4.061 13.939c-.293.293-.558.727-.75 1.188C3.119 15.59 3 16.086 3 16.5V21h4.5c.414 0 .908-.119 1.371-.311c.463-.192.896-.457 1.189-.75L21 9a1.497 1.497 0 0 0 0-2.121zM5.768 15.061l8.293-8.293L15.293 8L7 16.293l-1.232-1.232zM7.5 19H6l-1-1v-1.5c0-.077.033-.305.158-.605c.01-.02 2.967 2.938 2.967 2.938c-.322.134-.548.167-.625.167zm1.439-.768L7.707 17L16 8.707l1.232 1.232l-8.293 8.293zm9-9L14.767 6.06l1.293-1.293l3.17 3.172l-1.291 1.293z"/>`),
		g.Group(children),
	)
}