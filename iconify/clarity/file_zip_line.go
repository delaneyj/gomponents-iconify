package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZipLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 30.2V12l-8.1-7.9H7.8C6.8 4.1 6 4.9 6 6v24.2c0 1 .7 1.8 1.7 1.8h20.4c1 0 1.8-.7 1.8-1.7c.1 0 .1-.1.1-.1zM22 6.6l5.6 5.4H22V6.6zM28 30H7.9L8 6h12v8h8v16z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M12 24c0 1.7 1.3 3 3 3s3-1.3 3-3v-4h-6v4zm1.4 0v-2.6h3.2V24c.1.9-.6 1.7-1.5 1.7c-.9.1-1.7-.6-1.7-1.5V24z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18.2 9c0-.6-.4-1-1-1H15v2h2.2c.6 0 1-.4 1-1z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M12.7 10c-.6 0-1 .4-1 1s.4 1 1 1H15v-2h-2.3z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M17.2 14c.6 0 1-.4 1-1s-.4-1-1-1H15v2h2.2z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M11.7 15c0 .6.4 1 1 1H15v-2h-2.3c-.5 0-1 .4-1 1z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M17.2 18c.6 0 1-.4 1-1s-.4-1-1-1H15v2h2.2z" class="clr-i-outline clr-i-outline-path-7"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}