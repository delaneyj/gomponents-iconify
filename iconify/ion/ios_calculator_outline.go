package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosCalculatorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M144 160h224V96H144v64zm16-48h192v32H160v-32z" fill="currentColor"/><path d="M368 64H144c-17.7 0-32 14.3-32 32v320c0 17.7 14.3 32 32 32h224c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32zm16 352c0 8.8-7.2 16-16 16H144c-8.8 0-16-7.2-16-16V96c0-8.8 7.2-16 16-16h224c8.8 0 16 7.2 16 16v320z" fill="currentColor"/><path d="M160 240v-48h-16v64h64v-16h-16z" fill="currentColor"/><path d="M160 320v-48h-16v64h64v-16h-16z" fill="currentColor"/><path d="M160 400v-48h-16v64h64v-16h-16z" fill="currentColor"/><path d="M240 240v-48h-16v64h64v-16h-16z" fill="currentColor"/><path d="M320 240v-48h-16v64h64v-16h-16z" fill="currentColor"/><path d="M240 320v-48h-16v64h64v-16h-16z" fill="currentColor"/><path d="M240 400v-48h-16v64h64v-16h-16z" fill="currentColor"/><path d="M320 400V272h-16v144h64v-16h-16z" fill="currentColor"/>`),
		g.Group(children),
	)
}