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
        public uint NodeId { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class NodeRegisterAck
    {
        [global::ProtoBuf.ProtoMember(1)]
        public NodeType NodeType { get; set; }

        [global::ProtoBuf.ProtoMember(2)]
        public uint NodeId { get; set; }

    }

}

#pragma warning restore CS1591, CS0612, CS3021
