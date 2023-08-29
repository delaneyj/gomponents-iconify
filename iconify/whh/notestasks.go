package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notestasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V256q0-53 37.5-90.5T128 128h64V64q0-27 18.5-45.5T256 0t45.5 18.5T320 64v64h128V64q0-27 19-45.5T512.5 0t45 18.5T576 64v64h128V64q0-27 19-45.5T768.5 0t45 18.5T832 64v64h64q53 0 90.5 37.5T1024 256v640q0 53-37.5 90.5T896 1024zm0-768h-64q0 26-18.5 45t-45 19t-45.5-19t-19-45H576q0 26-18.5 45t-45 19t-45.5-19t-19-45H320q0 26-18.5 45t-45 19t-45.5-19t-19-45h-64v640h768V256zM362 529l86 86l213-214q19-18 44.5-18t44 18t18.5 44t-18 44L502 736q-6 11-8 13q-19 19-46 18q-28 1-47-18q-2-2-8-13L273 617q-18-18-18-44t18.5-44t44-18t44.5 18z"/>`),
		g.Group(children),
	)
}