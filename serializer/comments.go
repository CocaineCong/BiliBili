package serializer

import (
	"BiliBili.com/model"
)



func BuildComment(item model.CommentStrut) model.CommentStrut {
	return model.CommentStrut{
		ID:item.ID,
		CreatedAt  :item.CreatedAt,
		Content  :item.Content,
		Uid  :item.Uid,
		Name  :item.Name,
		Avatar  :item.Avatar,
		Reply  :BuildReplies(item.Reply),
	}
}


func BuildReply (item model.ReplyStrut) model.ReplyStrut {
	return model.ReplyStrut{
		ID:item.ID,
		CreatedAt  :item.CreatedAt,
		Content  :item.Content,
		Uid  :item.Uid,
		Name  :item.Name,
		Avatar  :item.Avatar,
		ReplyUid :item.ReplyUid,
		ReplyName:item.ReplyName,
	}
}

func BuildReplies(items []model.ReplyStrut) (replies []model.ReplyStrut){
	for _,item := range items {
		reply := BuildReply(item)
		replies = append(replies, reply)
	}
	return replies
}


func BuildComments(items []model.CommentStrut) (favors []model.CommentStrut) {
	for _,item := range items {
		favor := BuildComment(item)
		favors=append(favors, favor)
	}
	return favors
}

