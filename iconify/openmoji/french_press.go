package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrenchPress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M22 16h28v6H23l-5-4l4-2Z"/><path fill="#A57939" fill-rule="evenodd" d="M50 21v35a4 4 0 0 1-4 4H26a4 4 0 0 1-4-4V21h28Z" clip-rule="evenodd"/><path fill="#3F3F3F" d="M35.5 5a3.5 3.5 0 0 0-3.163 5h6.326A3.5 3.5 0 0 0 35.5 5Z"/><path fill="#9B9B9A" d="M34 15.5h3V26l-3-2v-8.5Z"/><path fill="#3F3F3F" fill-rule="evenodd" d="M50 16.5V15a3 3 0 0 0-3-3H25a3 3 0 0 0-3 3v1.5h28Z" clip-rule="evenodd"/><path fill="#6A462F" fill-rule="evenodd" d="M50 56H22a4 4 0 0 0 4 4h20a4 4 0 0 0 4-4Z" clip-rule="evenodd"/><path fill="#3F3F3F" fill-rule="evenodd" d="M50 46v4c7.732 0 14-6.268 14-14s-6.268-14-14-14v4c5.523 0 10 4.477 10 10s-4.477 10-10 10Z" clip-rule="evenodd"/><path fill="none" stroke="#3F3F3F" stroke-linecap="round" stroke-width="2" d="M35.5 5a3.5 3.5 0 0 0-3.163 5h6.326A3.5 3.5 0 0 0 35.5 5Z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M22 15a3 3 0 0 1 3-3h22a3 3 0 0 1 3 3v41a4 4 0 0 1-4 4H26a4 4 0 0 1-4-4V21.5L17.5 18l4.5-2.5V15Z" clip-rule="evenodd"/><path d="M50 50h-1v1h1v-1Zm0-4v-1h-1v1h1Zm0-24v-1h-1v1h1Zm0 4h-1v1h1v-1Zm1 24v-4h-2v4h2Zm12-14c0 7.18-5.82 13-13 13v2c8.284 0 15-6.716 15-15h-2ZM50 23c7.18 0 13 5.82 13 13h2c0-8.284-6.716-15-15-15v2Zm1 3v-4h-2v4h2Zm-1 1a9 9 0 0 1 9 9h2c0-6.075-4.925-11-11-11v2Zm9 9a9 9 0 0 1-9 9v2c6.075 0 11-4.925 11-11h-2Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M22 15.5h25M32.337 10a3.5 3.5 0 1 1 6.326 0M34 16v8"/><path stroke="#000" stroke-linecap="round" stroke-width="2" d="M37 16v10m-14-4h8"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M40 22h10"/>`),
		g.Group(children),
	)
}