package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.953 14H6.016v1.031H5V16h7v-.969h-1.047V14zM8.016 4v8.953h.953V4H13.5L8.969 2.833V1.016h-.953v1.817L3.469 4h4.547zM3.492 5.979l1.9 4.005l.566-.207l-2.225-4.609c-.053-.098-.157-.16-.298-.152c-.128.005-.239.075-.279.175L1.012 9.8l.588.162l1.892-3.983zm9.69-.784l-2.158 4.619l.592.162l1.902-3.99L15.43 10l.57-.208l-2.238-4.619c-.053-.097-.16-.159-.299-.151c-.13.003-.24.075-.281.173zm2.802 5.836c0 1.061-1.112 1.922-2.484 1.922c-1.372 0-2.484-.861-2.484-1.922h4.968zm-10 0c0 1.061-1.112 1.922-2.484 1.922c-1.372 0-2.484-.861-2.484-1.922h4.968z"/>`),
		g.Group(children),
	)
}