package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCopyLink0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M12 9.927V7a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v26a3 3 0 0 1-3 3h-2.983"/><rect width="34" height="34" x="4" y="10" fill="#555" stroke-linejoin="round" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m18.44 23.11l5.292-5.51c1.451-1.451 3.837-1.42 5.328.072c1.491 1.491 1.523 3.877.072 5.328l-1.91 2.023m-13.756 3.724c-.51.51-1.565 1.53-1.565 1.53c-1.452 1.451-1.492 4.038 0 5.53c1.49 1.49 3.876 1.523 5.328.071l5.164-4.688"/><path stroke-linecap="round" stroke-linejoin="round" d="M18.663 28.328a3.863 3.863 0 0 1-1.131-2.473A3.665 3.665 0 0 1 18.592 23m3.729 2.861c1.491 1.491 1.523 3.877.072 5.329"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCopyLink0)"/>`),
		g.Group(children),
	)
}