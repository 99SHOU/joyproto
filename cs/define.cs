// This file was generated by a tool; you should avoid making direct changes.
// Consider using 'partial classes' to extend these types
// Input: define.proto

#pragma warning disable CS1591, CS0612, CS3021

namespace Pb
{

    [global::ProtoBuf.ProtoContract()]
    public enum MsgID
    {
        [global::ProtoBuf.ProtoEnum(Name = @"FWM_NODE_REGISTER_REQ")]
        FwmNodeRegisterReq = 0,
        [global::ProtoBuf.ProtoEnum(Name = @"FWM_NODE_REGISTER_ACK")]
        FwmNodeRegisterAck = 1,
        [global::ProtoBuf.ProtoEnum(Name = @"FWM_SET_NODE_STATU")]
        FwmSetNodeStatu = 2,
        [global::ProtoBuf.ProtoEnum(Name = @"FWM_GAME_NODE_LIST_REQ")]
        FwmGameNodeListReq = 3,
        [global::ProtoBuf.ProtoEnum(Name = @"FWM_GAME_NODE_LIST_ACK")]
        FwmGameNodeListAck = 4,
        [global::ProtoBuf.ProtoEnum(Name = @"FWM_GAME_MSG_TRANSFER")]
        FwmGameMsgTransfer = 999,
    }

    [global::ProtoBuf.ProtoContract(Name = @"NODE_TYPE")]
    public enum NodeType
    {
        [global::ProtoBuf.ProtoEnum(Name = @"INVALID")]
        Invalid = 0,
        [global::ProtoBuf.ProtoEnum(Name = @"MACHINE")]
        Machine = 1,
        [global::ProtoBuf.ProtoEnum(Name = @"CENTER")]
        Center = 2,
        [global::ProtoBuf.ProtoEnum(Name = @"GAME")]
        Game = 3,
        [global::ProtoBuf.ProtoEnum(Name = @"GATE")]
        Gate = 4,
        [global::ProtoBuf.ProtoEnum(Name = @"LOGIN")]
        Login = 5,
        [global::ProtoBuf.ProtoEnum(Name = @"CLIENT")]
        Client = 6,
    }

    [global::ProtoBuf.ProtoContract(Name = @"NODE_STATU")]
    public enum NodeStatu
    {
        [global::ProtoBuf.ProtoEnum(Name = @"NOT_READY")]
        NotReady = 0,
        [global::ProtoBuf.ProtoEnum(Name = @"READY")]
        Ready = 1,
        [global::ProtoBuf.ProtoEnum(Name = @"OFF_LINE")]
        OffLine = 2,
    }

    [global::ProtoBuf.ProtoContract(Name = @"AGENT_INFO_KEY")]
    public enum AgentInfoKey
    {
        [global::ProtoBuf.ProtoEnum(Name = @"NODE_PORT")]
        NodePort = 0,
        [global::ProtoBuf.ProtoEnum(Name = @"ACCOUNT")]
        Account = 1,
        [global::ProtoBuf.ProtoEnum(Name = @"CHARACTER_ID")]
        CharacterId = 2,
        [global::ProtoBuf.ProtoEnum(Name = @"GAME_NODE_ID")]
        GameNodeId = 3,
    }

}

#pragma warning restore CS1591, CS0612, CS3021
