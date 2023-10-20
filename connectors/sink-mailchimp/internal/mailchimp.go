package internal

import (
	ce "github.com/cloudevents/sdk-go/v2"
	"github.com/hanzoai/gochimp3"
	"github.com/pkg/errors"
)

func (s *mailchimpSink) getList(event *ce.Event) (*gochimp3.ListResponse, error) {
	id, ok := event.Extensions()[AttrAudienceID].(string)
	if !ok {
		id = s.config.AudienceID
	}
	if list, exist := s.listMap[id]; exist {
		return list, nil
	}
	list, err := s.api.GetList(id, nil)
	if err != nil {
		return nil, err
	}
	s.listMap[id] = list
	return list, nil
}

func (s *mailchimpSink) addMember(list *gochimp3.ListResponse, event *ce.Event) error {
	member, err := event2Member(event)
	if err != nil {
		return err
	}
	// add subscriber
	_, err = list.CreateMember(&member.MemberRequest)
	if err != nil {
		return errors.Wrap(err, "add member error")
	}
	return nil
}

func (s *mailchimpSink) putMember(list *gochimp3.ListResponse, event *ce.Event) error {
	member, err := event2Member(event)
	if err != nil {
		return err
	}
	_, err = list.AddOrUpdateMember(emailHash(member.EmailAddress), &member.MemberRequest)
	if err != nil {
		return errors.Wrap(err, "put member error")
	}
	return nil
}

func (s *mailchimpSink) updateMember(list *gochimp3.ListResponse, event *ce.Event) error {
	member, err := event2Member(event)
	if err != nil {
		return err
	}
	_, err = list.UpdateMember(emailHash(member.EmailAddress), &member.MemberRequest)
	if err != nil {
		return errors.Wrap(err, "update member error")
	}
	return nil
}

func (s *mailchimpSink) archiveMember(list *gochimp3.ListResponse, event *ce.Event) error {
	member, err := event2Member(event)
	if err != nil {
		return err
	}
	_, err = list.DeleteMember(emailHash(member.EmailAddress))
	if err != nil {
		return errors.Wrap(err, "archive member error")
	}
	return nil
}

func (s *mailchimpSink) deleteMember(list *gochimp3.ListResponse, event *ce.Event) error {
	member, err := event2Member(event)
	if err != nil {
		return err
	}
	_, err = list.DeleteMemberPermanent(emailHash(member.EmailAddress))
	if err != nil {
		return errors.Wrap(err, "delete member error")
	}
	return nil
}
