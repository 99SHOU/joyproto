// This file was generated by a tool; you should avoid making direct changes.
// Consider using 'partial classes' to extend these types
// Input: framework_msg.proto

#pragma warning disable CS1591, CS0612, CS3021

namespace Pb
{

    [global::ProtoBuf.ProtoContract()]
    public partial class NodeRegisterReq
    {
        [global::ProtoBuf.ProtoMember(1)]
        public NodeType NodeType { get; set; }

        [global::ProtoBuf.ProtoMember(2)]
        public NodeStatu NodeStatu { get; set; }

        [global::ProtoBuf.ProtoMember(3)]
        public uint NodeId { get; set; }

        [global::ProtoBuf.ProtoMember(4)]
        public uint NodePort { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class NodeRegisterAck
    {
        [global::ProtoBuf.ProtoMember(1)]
        public NodeType NodeType { get; set; }

        [global::ProtoBuf.ProtoMember(2)]
        public NodeStatu NodeStatu { get; set; }

        [global::ProtoBuf.ProtoMember(3)]
        public uint NodeId { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class SetNodeStatu
    {
        [global::ProtoBuf.ProtoMember(1)]
        public NodeStatu NodeStatu { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class NodeInfo
    {
        [global::ProtoBuf.ProtoMember(1)]
        public NodeType NodeType { get; set; }

        [global::ProtoBuf.ProtoMember(2)]
        public NodeStatu NodeStatu { get; set; }

        [global::ProtoBuf.ProtoMember(3)]
        public uint NodeId { get; set; }

        [global::ProtoBuf.ProtoMember(4)]
        [global::System.ComponentModel.DefaultValue("")]
        public string Addr { get; set; } = "";

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class GameNodeListReq
    {
    }

    [global::ProtoBuf.ProtoContract()]
    public partial class GameNodeListAck
    {
        [global::ProtoBuf.ProtoMember(1)]
        public global::System.Collections.Generic.List<NodeInfo> NodeInfos { get; } = new global::System.Collections.Generic.List<NodeInfo>();

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class GameMsgTransfer
    {
        [global::ProtoBuf.ProtoMember(1)]
        public ulong CharacterId { get; set; }

        [global::ProtoBuf.ProtoMember(2)]
        public byte[] MsgId { get; set; }

        [global::ProtoBuf.ProtoMember(3)]
        public byte[] MsgBody { get; set; }

    }

}

#pragma warning restore CS1591, CS0612, CS3021
