package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorLoopCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="M20.248 6.75H8.5A1.5 1.5 0 0 0 7 8.25v7.651a1.5 1.5 0 0 0 1.5 1.5h11.748a1.5 1.5 0 0 0 1.5-1.5V8.25a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill-rule="evenodd" d="M6 8.25a2.5 2.5 0 0 1 2.5-2.5h11.748a2.5 2.5 0 0 1 2.5 2.5v7.651a2.5 2.5 0 0 1-2.5 2.5h-4.874v2.278a1 1 0 1 1-2 0V18.4H8.5A2.5 2.5 0 0 1 6 15.9V8.25Zm2.5-.5a.5.5 0 0 0-.5.5v7.651a.5.5 0 0 0 .5.5h11.748a.5.5 0 0 0 .5-.5V8.25a.5.5 0 0 0-.5-.5H8.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.687 21.34a1 1 0 0 1 1-1h7.374a1 1 0 1 1 0 2h-7.374a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M11.727 16.566a3.193 3.193 0 1 1-5.639-2.997a3.193 3.193 0 0 1 5.639 2.997Z"/><path fill-rule="evenodd" d="M9.937 13.132a2.193 2.193 0 1 0-2.058 3.872a2.193 2.193 0 0 0 2.058-3.873ZM5.205 13.1a4.193 4.193 0 1 1 2.674 6.034l-1.677 3.155a1 1 0 1 1-1.766-.939l1.677-3.155a4.194 4.194 0 0 1-.908-5.095Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M5 7.5a2 2 0 0 1 2-2h11.748a2 2 0 0 1 2 2v7.651a2 2 0 0 1-2 2h-5.374v2.778a.5.5 0 0 1-1 0V17.15H11.5a.5.5 0 0 1 0-1h7.248a1 1 0 0 0 1-1V7.5a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1V10a.5.5 0 0 1-1 0V7.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.687 20.59a.5.5 0 0 1 .5-.5h7.374a.5.5 0 1 1 0 1H9.187a.5.5 0 0 1-.5-.5Zm-.016-8.65a2.693 2.693 0 1 0-2.527 4.756A2.693 2.693 0 0 0 8.67 11.94Zm-4.524.645a3.693 3.693 0 1 1 1.985 5.199l-1.871 3.52a.5.5 0 0 1-.883-.47l1.87-3.52a3.694 3.694 0 0 1-1.101-4.73Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}