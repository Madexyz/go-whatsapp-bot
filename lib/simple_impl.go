package lib

import (
	"context"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

type SimpleImpl struct {
	VClient *whatsmeow.Client
	Msg     *events.Message
}

func NewSimpleImpl(Cli *whatsmeow.Client, m *events.Message) *SimpleImpl {
	return &SimpleImpl{
		VClient: Cli,
		Msg:     m,
	}
}

func (simp *SimpleImpl) Reply(teks string) {
	simp.VClient.SendMessage(context.Background(), simp.Msg.Info.Chat, "", &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: proto.String(teks),
			ContextInfo: &waProto.ContextInfo{
				StanzaId:      &simp.Msg.Info.ID,
				Participant:   proto.String(simp.Msg.Info.Sender.String()),
				QuotedMessage: simp.Msg.Message,
			},
		},
	})
}

func (simp *SimpleImpl) SendHydratedBtn(jid types.JID, teks string, foter string, buttons []*waProto.HydratedTemplateButton) {
	simp.VClient.SendMessage(context.Background(), jid, "", &waProto.Message{
		TemplateMessage: &waProto.TemplateMessage{
			HydratedTemplate: &waProto.HydratedFourRowTemplate{
				HydratedContentText: proto.String(teks),
				HydratedFooterText:  proto.String(foter),
				HydratedButtons:     buttons,
			},
		},
	})
}

func (simp *SimpleImpl) GetCMD() string {
	extended := simp.Msg.Message.GetExtendedTextMessage().GetText()
	text := simp.Msg.Message.GetConversation()
	imageMatch := simp.Msg.Message.GetImageMessage().GetCaption()
	videoMatch := simp.Msg.Message.GetVideoMessage().GetCaption()
	var command string
	if text != "" {
		command = text
	} else if imageMatch != "" {
		command = imageMatch
	} else if videoMatch != "" {
		command = videoMatch
	} else if extended != "" {
		command = extended
	}
	return command
}
