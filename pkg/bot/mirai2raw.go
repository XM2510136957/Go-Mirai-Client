package bot

import (
	"fmt"
	"html"

	"github.com/ProtobufBot/Go-Mirai-Client/pkg/clz"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

func MiraiMsgToRawMsg(cli *client.QQClient, messageChain []message.IMessageElement) string {
	result := ""
	for _, element := range messageChain {
		switch elem := element.(type) {
		case *message.TextElement:
			result += elem.Content
		case *message.AtElement:
			result += fmt.Sprintf(`<at qq="%d"/>`, elem.Target)
		case *message.DiceElement:
			result += fmt.Sprintf(`<dice value="%d"/>`, elem.Value)
		case *message.FingerGuessingElement:
			result += fmt.Sprintf(`<finger_guessing value="%d" name="%s"/>`, elem.Value, elem.Name)
		case *message.FriendImageElement:
			result += fmt.Sprintf(`<image image_id="%s" url="%s"/>`, html.EscapeString(elem.ImageId), html.EscapeString(elem.Url))
		case *message.GroupImageElement:
			result += fmt.Sprintf(`<image image_id="%s" url="%s"/>`, html.EscapeString(elem.ImageId), html.EscapeString(elem.Url))
		case *clz.LocalImageElement:
			result += fmt.Sprintf(`<image url="%s"/>`, html.EscapeString(elem.Url))
		case *message.FaceElement:
			result += fmt.Sprintf(`<face id="%d" name="%s"/>`, elem.Index, html.EscapeString(elem.Name))
		case *message.VoiceElement:
			result += fmt.Sprintf(`<voice url="%s"/>`, html.EscapeString(elem.Url))
		case *message.ServiceElement:
			result += fmt.Sprintf(`<service id="%d" content="%s" res_id="%s" sub_type="%s"/>`, elem.Id, html.EscapeString(elem.Content), elem.ResId, elem.SubType)
		case *message.LightAppElement:
			result += fmt.Sprintf(`<light_app content="%s"/>`, html.EscapeString(elem.Content))
		case *message.ShortVideoElement:
			result += fmt.Sprintf(`<video name="%s" url="%s"/>`, html.EscapeString(elem.Name), html.EscapeString(cli.GetShortVideoUrl(elem.Uuid, elem.Md5)))
		case *message.ReplyElement:
			result += fmt.Sprintf(`<reply time="%d" sender="%d" raw_message="%s" reply_seq="%d"/>`, elem.Time, elem.Sender, html.EscapeString(MiraiMsgToRawMsg(cli, elem.Elements)), elem.ReplySeq)
		case *clz.MyVideoElement:
			result += fmt.Sprintf(`<video url="%s" cover="%s"/>`, html.EscapeString(elem.Url), html.EscapeString(elem.CoverUrl))
		}
	}

	return result
}
