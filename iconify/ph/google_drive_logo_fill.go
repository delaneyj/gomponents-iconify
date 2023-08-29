package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M237.6 151.78L169.13 39.52A15.91 15.91 0 0 0 155.56 32h-55.13a15.89 15.89 0 0 0-13.56 7.52l-.05.07L18.44 151.7a16 16 0 0 0-.33 16.42l27.32 47.82A16 16 0 0 0 59.32 224h137.35a16 16 0 0 0 13.89-8.06l27.32-47.82a15.91 15.91 0 0 0-.28-16.34ZM219 152h-46.48l-35.19-58.67l22.75-37.92ZM92.53 168h70.94l24 40H68.53Zm9.6-16L128 108.88L153.87 152Zm-6.22-96.59l22.76 37.92L83.47 152H37Z"/>`),
		g.Group(children),
	)
}