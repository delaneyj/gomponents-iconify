package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vigo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVigo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="3.833"><path fill="#555" d="M24 44.125c9.527 0 17.25-7.616 17.25-17.01c0-2.306-.315-4.5-.99-6.591c-.39-1.207-.704-1.783-1.167-2.894c-.463-1.111-1.079-2.05-1.78-2.993c-.7-.943-2.067-2.798-3.012-3.17c0 1.461-2.873 6.568-3.975 6.824c-1.103.257.815-5.369-1.628-9.217c-2.443-3.848-6.393-6.374-6.393-4.644c0 1.729-1.093 4.632-2.194 6.072c-1.1 1.439-2.886 2.675-4.13 3.256c-1.243.581-.793-2.908-1.726-2.292c-.94.621-2.278 2.351-2.98 3.437c-2.462 3.801-4.525 7.515-4.525 12.212c0 9.394 7.723 17.01 17.25 17.01Z"/><path fill="#fff" d="M19.296 23.733v8.737a1.878 1.878 0 0 0 2.863 1.599l8.385-5.169a.939.939 0 0 0 0-1.598l-8.385-5.168a1.878 1.878 0 0 0-2.863 1.599Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVigo0)"/>`),
		g.Group(children),
	)
}