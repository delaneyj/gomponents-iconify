package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftFacingFistMediumDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiLeftFacingFistMediumDarkSkinTone0" d="M15.31 29.17c0 4.062 4.892 4.344 10.94 4.344c-.5 4.031-2.594 8.625-10.03 13"/></defs><path fill="#a57939" d="m63.858 33.942l-8.881 1.886c-1.051.246-2.034.55-3.291 2.112l-8.311 9.648c-.776.933-1.248 1.483-2.496 1.876l-20.532 7.604c-4.362 1.11-8.291 0-7.309-8.282c-4.263 1.621-8.32.305-8.026-3.399l-.01-20.1c.57-4.185 4.313-6.18 7.545-6.513l50.427-2.044c5 4.51 5.393 12.212.884 17.212z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2.036" transform="matrix(.9826 0 0 .9823 .266 1.018)"><path d="m43.88 47.41l8.469-9.812c1.281-1.594 2.281-1.906 3.344-2.156l9.04-1.92c4.59-5.084 4.19-12.93-.894-17.52l-51.33 2.083c-3.292.333-7.1 2.364-7.682 6.622l.016 20.46c-.303 3.771 3.823 5.115 8.167 3.458c-1 8.438 3 9.562 7.438 8.438l20.91-7.739"/><use href="#openmojiLeftFacingFistMediumDarkSkinTone0"/><path d="m43.88 47.41l8.469-9.812c1.281-1.594 2.281-1.906 3.344-2.156l9.04-1.92c4.59-5.084 4.19-12.93-.894-17.52l-51.33 2.083c-3.292.333-7.1 2.364-7.682 6.622l.016 20.46c-.303 3.771 3.823 5.115 8.167 3.458c-1 8.438 3 9.562 7.438 8.438l20.91-7.739c1.296-.68 1.853-1.293 2.53-1.917z"/><use href="#openmojiLeftFacingFistMediumDarkSkinTone0"/></g>`),
		g.Group(children),
	)
}