package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.905 14.325l-1.83-10.04a1.507 1.507 0 0 0-1.47-1.22h-11.2A1.493 1.493 0 0 0 4.925 4.3l-1.84 10.03a2.452 2.452 0 0 0-.02.27v4.84a1.5 1.5 0 0 0 1.5 1.5h14.87a1.511 1.511 0 0 0 1.5-1.5V14.6a1.241 1.241 0 0 0-.03-.275Zm-15-9.85a.5.5 0 0 1 .5-.41h11.2a.511.511 0 0 1 .49.4l1.74 9.54H4.165Zm14.03 14.96a.5.5 0 0 1-.5.5H4.565a.5.5 0 0 1-.5-.5l.01-4.43h15.86Z"/><circle cx="5.561" cy="17.47" r=".5" fill="currentColor"/><circle cx="7.561" cy="17.47" r=".5" fill="currentColor"/><path fill="currentColor" d="M18.45 17.97a.5.5 0 0 0 0-1h-5a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}