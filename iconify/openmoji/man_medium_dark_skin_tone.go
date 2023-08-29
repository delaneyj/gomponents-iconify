package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManMediumDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiManMediumDarkSkinTone0" d="M41.9 30c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2m-8 0c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2"/></defs><path fill="#92D3F5" d="M17 61v-4c0-5 5-9 10-9c6 5 12 5 18 0c5 0 10 4 10 9v4"/><path d="M26 38c-3 0-4-7-4-14c0-6 5-12 14-12s14 6 14 12c0 7-1 14-4 14"/><path fill="#a57939" d="M24.9 31c-.1 8 4.1 14 11 14C43 45 47 39 47 31c0-5-3-10-3-10c-8 0-10 3-16 1c0 0-3 4-3.1 9z"/><use href="#openmojiManMediumDarkSkinTone0"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M32.9 38.2c2.1.8 4.1.8 6 0M26 38c-3 0-4-7-4-14c0-6 5-12 14-12s14 6 14 12c0 7-1 14-4 14M17 60v-3c0-5 5-9 10-9c6 5 12 5 18 0c5 0 10 4 10 9v3"/><use href="#openmojiManMediumDarkSkinTone0"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M24.9 31c-.1 8 4.9 14 11 14c6 0 11.1-6 11.1-14c0-5-3-11-3-11c-8 0-10 3-16 1c0 0-3 5-3.1 10z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M32.9 38.2c2.1.8 4.1.8 6 0"/>`),
		g.Group(children),
	)
}