kd                   = require 'kd'
React                = require 'kd-react'
immutable            = require 'immutable'
ActivityFlux         = require 'activity/flux'
ChatPane             = require 'activity/components/chatpane'
ChatInputWidget      = require 'activity/components/chatinputwidget'
ChatPaneWrapperMixin = require 'activity/components/chatpane/chatpanewrappermixin'
FollowChannelBox     = require 'activity/components/followchannelbox'

{ message: messageActions, command: commandActions } = ActivityFlux.actions

module.exports = class PublicChatPane extends React.Component

  @defaultProps =
    thread   : immutable.Map()

  channel: (keyPath...) -> @props.thread?.getIn ['channel'].concat keyPath


  onSubmit: ({ value }) ->

    return  unless value

    messageActions.createMessage @channel('id'), value


  onCommand: ({ command }) -> commandActions.executeCommand command, @channel()


  onLoadMore: ->

    return  unless (messages = @props.thread.get 'messages').size

    messageActions.loadMessages @channel('id'),
      from: messages.first().get 'createdAt'


  onInviteOthers: ->

    return  unless input = @refs.chatInputWidget

    input.setCommand '/invite @'


  renderFooter: ->

    return null  unless @props.thread?.get 'messages'

    isParticipant = @channel 'isParticipant'
    isPrivate = 'privatemessage' is @channel 'typeConstant'

    disabledFeatures = []
    disabledFeatures = disabledFeatures.concat ['search']  if isPrivate

    <footer className="PublicChatPane-footer ChatPaneFooter">
      <ChatInputWidget
        ref='chatInputWidget'
        className={unless isParticipant then 'hidden'}
        onSubmit={@bound 'onSubmit'}
        onCommand={@bound 'onCommand'}
        channelId={@channel 'id'}
        onResize={@bound 'onResize'}
        disabledFeatures={disabledFeatures} />
      <FollowChannelBox
        className={if isParticipant then 'hidden'}
        thread={@props.thread} />
    </footer>


  render: ->

    return null  unless @props.thread

    <div>
      <ChatPane
        ref='chatPane'
        key={@props.thread.get 'channelId'}
        thread={@props.thread}
        className='PublicChatPane'
        onSubmit={@bound 'onSubmit'}
        onLoadMore={@bound 'onLoadMore'}
        onInviteOthers={@bound 'onInviteOthers'}
      />
      {@renderFooter()}
    </div>


PublicChatPane.include [ChatPaneWrapperMixin]

